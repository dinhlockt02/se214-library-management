package mysql

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	coreerror "daijoubuteam.xyz/se214-library-management/core/error"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"time"
)

type PhieuMuonRepository struct {
	*sqlx.DB
}

func (r PhieuMuonRepository) GetDanhSachPhieuMuon() (_ []*entity.PhieuMuon, err error) {
	tx := r.DB.MustBegin()
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	var rows *sqlx.Rows

	rows, err = tx.Queryx(
		`SELECT MaPhieuMuon, MaDocGia, MaSach, NgayMuon FROM PhieuMuon WHERE MaPhieuMuon NOT IN (
    			SELECT MaPhieuMuon FROM PhieuTra
			)`,
	)
	defer rows.Close()
	var danhSachPhieuMuon []*entity.PhieuMuon
	for rows.Next() {
		var mpm, mdg, ms string
		var ngayMuon time.Time
		err = rows.Scan(&mpm, &mdg, &ms, &ngayMuon)
		if err != nil {
			return nil, coreerror.NewInternalServerError("database error: can't not query phieu muon", err)
		}
		var maPhieuMuon, maSach *entity.ID
		if maPhieuMuon, err = entity.StringToID(mpm); err != nil {
			return nil, coreerror.NewInternalServerError("database error: can't not query phieu muon", err)
		}
		if maSach, err = entity.StringToID(ms); err != nil {
			return nil, coreerror.NewInternalServerError("database error: can't not query phieu muon", err)
		}
		phieuMuon := entity.NewPhieuMuon(&entity.DocGia{MaDocGia: mdg}, &ngayMuon, &entity.Sach{MaSach: maSach})
		phieuMuon.MaPhieuMuon = maPhieuMuon
		danhSachPhieuMuon = append(danhSachPhieuMuon, phieuMuon)
	}
	for i, _ := range danhSachPhieuMuon {
		if danhSachPhieuMuon[i].Sach, err = NewSachRepository(r.DB).getSachWithTx(danhSachPhieuMuon[i].Sach.MaSach, tx); err != nil {
			return nil, err
		}
		if danhSachPhieuMuon[i].DocGia, err = NewDocGiaRepository(r.DB).getDocGiaWithTx(tx, danhSachPhieuMuon[i].DocGia.MaDocGia); err != nil {
			return nil, err
		}
	}
	return danhSachPhieuMuon, nil
}

func (r PhieuMuonRepository) GetPhieuMuonByDocGia(maDocGia string) (_ []*entity.PhieuMuon, err error) {
	tx := r.DB.MustBegin()
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	var rows *sqlx.Rows

	rows, err = tx.Queryx(
		`SELECT MaPhieuMuon, MaDocGia, MaSach, NgayMuon FROM PhieuMuon WHERE MaPhieuMuon NOT IN (
    			SELECT MaPhieuMuon FROM PhieuTra
			) AND MaDocGia = ?`,
		maDocGia,
	)
	defer rows.Close()
	var danhSachPhieuMuon []*entity.PhieuMuon
	for rows.Next() {
		var mpm, mdg, ms string
		var ngayMuon time.Time
		err = rows.Scan(&mpm, &mdg, &ms, &ngayMuon)
		if err != nil {
			return nil, coreerror.NewInternalServerError("database error: can't not query phieu muon", err)
		}
		var maPhieuMuon, maSach *entity.ID
		if maPhieuMuon, err = entity.StringToID(mpm); err != nil {
			return nil, coreerror.NewInternalServerError("database error: can't not query phieu muon", err)
		}
		if maSach, err = entity.StringToID(ms); err != nil {
			return nil, coreerror.NewInternalServerError("database error: can't not query phieu muon", err)
		}
		phieuMuon := entity.NewPhieuMuon(&entity.DocGia{MaDocGia: maDocGia}, &ngayMuon, &entity.Sach{MaSach: maSach})
		phieuMuon.MaPhieuMuon = maPhieuMuon
		danhSachPhieuMuon = append(danhSachPhieuMuon, phieuMuon)
	}
	for i, _ := range danhSachPhieuMuon {
		if danhSachPhieuMuon[i].Sach, err = NewSachRepository(r.DB).getSachWithTx(danhSachPhieuMuon[i].Sach.MaSach, tx); err != nil {
			return nil, err
		}
		if danhSachPhieuMuon[i].DocGia, err = NewDocGiaRepository(r.DB).getDocGiaWithTx(tx, danhSachPhieuMuon[i].DocGia.MaDocGia); err != nil {
			return nil, err
		}
	}
	return danhSachPhieuMuon, nil
}

func (r PhieuMuonRepository) CreatePhieuMuon(phieuMuon *entity.PhieuMuon) (_ *entity.PhieuMuon, err error) {
	tx := r.DB.MustBegin()
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	_, err = tx.Exec(
		`INSERT INTO PhieuMuon(MaPhieuMuon, MaDocGia, MaSach, NgayMuon) VALUES (?,?,?,?)`,
		phieuMuon.MaPhieuMuon.String(), phieuMuon.MaDocGia, phieuMuon.MaSach, phieuMuon.NgayMuon,
	)
	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: create phieu muon failed", err)
	}
	_, err = tx.Exec(
		`UPDATE Sach SET TinhTrang = FALSE WHERE MaSach = ?`,
		phieuMuon.MaSach.String(),
	)
	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: create phieu muon failed", err)
	}

	return phieuMuon, nil
}

func (r PhieuMuonRepository) GetPhieuMuonByMaSach(maSach *entity.ID) (_ *entity.PhieuMuon, err error) {
	tx := r.DB.MustBegin()
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	row := tx.QueryRowx(
		`SELECT MaPhieuMuon, MaDocGia, NgayMuon FROM PhieuMuon WHERE MaPhieuMuon NOT IN (
    			SELECT MaPhieuMuon FROM PhieuTra
			) AND MaSach = ?`,
		maSach.String(),
	)
	var mpm, mdg string
	var ngayMuon time.Time
	err = row.Scan(&mpm, &mdg, &ngayMuon)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, coreerror.NewNotFoundError("phieu muon not found", err)
		}
		return nil, coreerror.NewInternalServerError("database error: can't not query phieu muon", err)
	}
	var maPhieuMuon *entity.ID
	if maPhieuMuon, err = entity.StringToID(mpm); err != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not query phieu muon", err)
	}
	phieuMuon := entity.NewPhieuMuon(&entity.DocGia{MaDocGia: mdg}, &ngayMuon, &entity.Sach{MaSach: maSach})
	phieuMuon.MaPhieuMuon = maPhieuMuon
	if phieuMuon.Sach, err = NewSachRepository(r.DB).getSachWithTx(phieuMuon.Sach.MaSach, tx); err != nil {
		return nil, err
	}
	if phieuMuon.DocGia, err = NewDocGiaRepository(r.DB).getDocGiaWithTx(tx, mdg); err != nil {
		return nil, err
	}
	return phieuMuon, nil
}

func NewPhieuMuonRepository(db *sqlx.DB) PhieuMuonRepository {
	return PhieuMuonRepository{db}
}
