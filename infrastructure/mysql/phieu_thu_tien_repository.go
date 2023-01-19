package mysql

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	coreerror "daijoubuteam.xyz/se214-library-management/core/error"
	"github.com/jmoiron/sqlx"
)

type PhieuThuTienRepository struct {
	*sqlx.DB
}

func NewPhieuThuTienRepository(db *sqlx.DB) PhieuThuTienRepository {
	return PhieuThuTienRepository{
		db,
	}
}

func (r PhieuThuTienRepository) GetPhieuThuTienByMaDocGia(maDocGia string) (_ []*entity.PhieuThuTien, err error) {
	tx := r.DB.MustBegin()
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	var rows *sqlx.Rows
	if rows, err = tx.Queryx(`SELECT MaPhieuThuTienPhat, TongNo, SoTienThu, ConLai, NgayThu FROM PhieuThuTienPhat WHERE MaDocGia = ?`, maDocGia); err != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not query phieu thu tien", err)
	}
	var danhSachPhieuThu []*entity.PhieuThuTien
	for rows.Next() {
		var phieuThu = &entity.PhieuThuTien{DocGia: &entity.DocGia{MaDocGia: maDocGia}}
		var mpt string
		if err = rows.Scan(&mpt, &(phieuThu.TongNo), &(phieuThu.SoTienThu), &(phieuThu.ConLai), &(phieuThu.NgayThu)); err != nil {
			return nil, coreerror.NewInternalServerError("database error: can't not query phieu thu tien", err)
		}
		if phieuThu.MaPhieuThu, err = entity.StringToID(mpt); err != nil {
			return nil, coreerror.NewInternalServerError("database error: can't not query phieu thu tien", err)
		}
		danhSachPhieuThu = append(danhSachPhieuThu, phieuThu)
	}
	for i, _ := range danhSachPhieuThu {
		danhSachPhieuThu[i].DocGia, err = NewDocGiaRepository(r.DB).getDocGiaWithTx(tx, maDocGia)
		if err != nil {
			return nil, coreerror.NewInternalServerError("database error: can't not query phieu thu tien", err)
		}
	}
	return danhSachPhieuThu, nil

}

func (r PhieuThuTienRepository) CreatePhieuThuTien(phieuThuTien *entity.PhieuThuTien) (_ *entity.PhieuThuTien, err error) {
	tx := r.DB.MustBegin()
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	_, err = tx.Exec(
		`INSERT INTO PhieuThuTienPhat(
                             MaPhieuThuTienPhat, 
                             MaDocGia, 
                             TongNo, 
                             SoTienThu, 
                             ConLai, 
                             NgayThu) 
				VALUES (?, ?, ? , ? , ? , ?)`,
		phieuThuTien.MaPhieuThu,
		phieuThuTien.DocGia.MaDocGia,
		phieuThuTien.TongNo,
		phieuThuTien.SoTienThu,
		phieuThuTien.ConLai,
		phieuThuTien.NgayThu)
	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not create phieu thu tien phat", err)
	}

	_, err = tx.Exec(
		`UPDATE DocGia SET TongNo = ? WHERE MaDocGia = ?`,
		phieuThuTien.DocGia.TongNo,
		phieuThuTien.DocGia.MaDocGia)
	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not create phieu thu tien phat", err)
	}

	return phieuThuTien, nil
}
