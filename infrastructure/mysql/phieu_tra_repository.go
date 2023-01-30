package mysql

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	coreerror "daijoubuteam.xyz/se214-library-management/core/error"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type PhieuTraRepository struct {
	*sqlx.DB
}

func NewPhieuTraRepository(db *sqlx.DB) PhieuTraRepository {
	return PhieuTraRepository{db}
}

func (r PhieuTraRepository) GetDanhSachPhieuTra() (_ []*entity.PhieuTra, err error) {
	tx := r.MustBegin()
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()
	var rows *sqlx.Rows
	if rows, err = tx.Queryx(`SELECT PM.MaPhieuMuon, TienPhat, NgayTra, GhiChu, MaDocGia, MaSach FROM PhieuTra PT JOIN PhieuMuon PM on PM.MaPhieuMuon = PT.MaPhieuMuon`); err != nil {
		return nil, coreerror.NewInternalServerError("databaser error: can't not query database", err)
	}
	defer rows.Close()
	var danhSachPhieuTra []*entity.PhieuTra
	for rows.Next() {
		var phieuTra = &entity.PhieuTra{PhieuMuon: &entity.PhieuMuon{DocGia: &entity.DocGia{}, Sach: &entity.Sach{}}}
		var mpm string
		var ms string
		var mdg string
		if err = rows.Scan(&mpm, &(phieuTra.TienPhat), &(phieuTra.NgayTra), &(phieuTra.GhiChu), &mdg, &ms); err != nil {
			return nil, coreerror.NewInternalServerError("databaser error: can't not query database", err)
		}
		var maPhieuMuon *entity.ID
		if maPhieuMuon, err = entity.StringToID(mpm); err != nil {
			return nil, coreerror.NewInternalServerError("databaser error: can't not query database", err)
		}

		var maSach *entity.ID
		if maSach, err = entity.StringToID(ms); err != nil {
			return nil, coreerror.NewInternalServerError("databaser error: can't not query database", err)
		}

		phieuTra.MaPhieuMuon = maPhieuMuon
		phieuTra.PhieuMuon.DocGia.MaDocGia = mdg
		phieuTra.MaSach = maSach

		danhSachPhieuTra = append(danhSachPhieuTra, phieuTra)

	}
	for i, _ := range danhSachPhieuTra {
		if danhSachPhieuTra[i].Sach, err = NewSachRepository(r.DB).getSachWithTx(danhSachPhieuTra[i].MaSach, tx); err != nil {
			return nil, err
		}
		if danhSachPhieuTra[i].DocGia, err = NewDocGiaRepository(r.DB).getDocGiaWithTx(tx, danhSachPhieuTra[i].MaDocGia); err != nil {
			return nil, err
		}
	}
	return danhSachPhieuTra, nil
}

func (r PhieuTraRepository) GetPhieuTraByDocGia(maDocGia string) (_ []*entity.PhieuTra, err error) {
	tx := r.MustBegin()
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()
	var rows *sqlx.Rows
	if rows, err = tx.Queryx(
		`SELECT PM.MaPhieuMuon, TienPhat, NgayTra, GhiChu, MaDocGia, MaSach
				FROM PhieuTra PT JOIN PhieuMuon PM ON PM.MaPhieuMuon = PT.MaPhieuMuon
				WHERE MaDocGia = ?`, maDocGia); err != nil {
		return nil, coreerror.NewInternalServerError("databaser error: can't not query database", err)
	}
	defer rows.Close()
	var danhSachPhieuTra []*entity.PhieuTra
	for rows.Next() {
		var phieuTra = &entity.PhieuTra{PhieuMuon: &entity.PhieuMuon{}}
		var mpm string
		var ms string
		var mdg string
		if err = rows.Scan(&mpm, &(phieuTra.TienPhat), &(phieuTra.NgayTra), &(phieuTra.GhiChu), &mdg, &ms); err != nil {
			return nil, coreerror.NewInternalServerError("databaser error: can't not query database", err)
		}
		var maPhieuMuon *entity.ID
		if maPhieuMuon, err = entity.StringToID(mpm); err != nil {
			return nil, coreerror.NewInternalServerError("databaser error: can't not query database", err)
		}
		var maSach *entity.ID
		if maSach, err = entity.StringToID(ms); err != nil {
			return nil, coreerror.NewInternalServerError("databaser error: can't not query database", err)
		}
		phieuTra.MaPhieuMuon = maPhieuMuon
		phieuTra.MaDocGia = maDocGia
		phieuTra.MaSach = maSach
		danhSachPhieuTra = append(danhSachPhieuTra, phieuTra)
	}
	for i, _ := range danhSachPhieuTra {
		if danhSachPhieuTra[i].Sach, err = NewSachRepository(r.DB).getSachWithTx(danhSachPhieuTra[i].MaSach, tx); err != nil {
			return nil, err
		}
		if danhSachPhieuTra[i].DocGia, err = NewDocGiaRepository(r.DB).getDocGiaWithTx(tx, danhSachPhieuTra[i].MaDocGia); err != nil {
			return nil, err
		}
	}
	return danhSachPhieuTra, nil
}

func (r PhieuTraRepository) CreatePhieuTra(phieuTra *entity.PhieuTra) (_ *entity.PhieuTra, err error) {
	tx := r.MustBegin()
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	_, err = tx.Exec(`INSERT INTO PhieuTra(MaPhieuMuon, TienPhat, NgayTra, GhiChu) VALUES (?, ?, ?, ?)`, phieuTra.MaPhieuMuon, phieuTra.TienPhat, phieuTra.NgayTra, phieuTra.GhiChu)
	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not insert into database", err)
	}
	_, err = tx.Exec(`UPDATE Sach SET TinhTrang = TRUE WHERE MaSach = ?`, phieuTra.MaSach.String())
	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not update database", err)
	}
	if phieuTra.TienPhat > 0 {
		_, err = tx.Exec(`UPDATE DocGia SET TongNo = ? WHERE MaDocGia = ?`, phieuTra.DocGia.TongNo, phieuTra.MaDocGia)
		if err != nil {
			fmt.Println(err)
			return nil, coreerror.NewInternalServerError("database error: can't not update database", err)
		}
	}
	return phieuTra, nil
}
