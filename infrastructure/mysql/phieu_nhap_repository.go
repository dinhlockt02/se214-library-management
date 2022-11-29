package mysql

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	coreerror "daijoubuteam.xyz/se214-library-management/core/error"
	"daijoubuteam.xyz/se214-library-management/core/repository"
	"daijoubuteam.xyz/se214-library-management/utils"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"reflect"
)

type PhieuNhapRepository struct {
	db                *sqlx.DB
	dauSachRepository repository.DauSachRepository
}

func NewPhieuNhapRepository(db *sqlx.DB, dauSachRepository repository.DauSachRepository) *PhieuNhapRepository {
	return &PhieuNhapRepository{
		db:                db,
		dauSachRepository: dauSachRepository,
	}
}

func (repo *PhieuNhapRepository) GetDanhSachPhieuNhap() (_ []*entity.PhieuNhap, err error) {
	tx := repo.db.MustBegin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// Query phieu nhap

	stmt, err := tx.Prepare(`SELECT MaPhieuNhap, NgayNhap, TongTien FROM PhieuNhap`)
	if err != nil {
		return nil, coreerror.NewInternalServerError("database err: can't not prepare query", err)
	}
	rows, err := stmt.Query()
	defer func() {
		rows.Close()
	}()
	if err != nil {
		return nil, coreerror.NewInternalServerError("database err: can't not execute query", err)
	}
	danhSachPhieuNhap := make([]*entity.PhieuNhap, 0)
	for rows.Next() {
		var phieuNhap *entity.PhieuNhap = &entity.PhieuNhap{}
		var maPhieuNhap string
		s := reflect.ValueOf(phieuNhap).Elem()
		numCols := s.NumField()
		columns := make([]interface{}, numCols)
		for i := 1; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}
		columns[0] = &maPhieuNhap
		err := rows.Scan(columns[0 : len(columns)-1]...)
		if err != nil {
			return nil, coreerror.NewInternalServerError("database err: can't not scan rows", err)
		}
		phieuNhap.MaPhieuNhap, err = entity.StringToID(maPhieuNhap)
		if err != nil {
			return nil, coreerror.NewInternalServerError("database err: can't not convert to id", err)
		}
		ctPhieuNhap, err := repo.getCtPhieuNhapByPhieuNhap(maPhieuNhap)
		if err != nil {
			return nil, err
		}
		phieuNhap.CtPhieuNhap = ctPhieuNhap
		danhSachPhieuNhap = append(danhSachPhieuNhap, phieuNhap)
	}
	return danhSachPhieuNhap, nil
}

func (repo *PhieuNhapRepository) getCtPhieuNhapByPhieuNhap(maPhieuNhap string) (_ []*entity.CtPhieuNhap, err error) {
	tx := repo.db.MustBegin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	stmt, err := tx.Prepare(`SELECT Sach.MaSach,MaDauSach, NhaXuatBan, TriGia, NamXuatBan, TinhTrang, DonGia FROM Sach INNER JOIN Ct_PhieuNhap CPN on Sach.MaSach = CPN.MaSach WHERE MaPhieuNhap = ?`)
	if err != nil {
		return nil, coreerror.NewInternalServerError("database err: can't not prepare query", err)
	}
	rows, err := stmt.Query(maPhieuNhap)
	defer func() {
		rows.Close()
	}()
	danhSachCtPhieuNhap := make([]*entity.CtPhieuNhap, 0)
	for rows.Next() {
		var MaSachDB string
		var MaDauSachDB string
		var NhaXuatBan string
		var TriGia uint
		var NamXuatBan uint
		var TinhTrang bool
		var DonGia uint
		err := rows.Scan(&MaSachDB, &MaDauSachDB, &NhaXuatBan, &TriGia, &NamXuatBan, &TinhTrang, &DonGia)
		if err != nil {
			return nil, coreerror.NewInternalServerError("database err: can't not scan rows", err)
		}

		MaSach, err := entity.StringToID(MaSachDB)
		if err != nil {
			return nil, coreerror.NewInternalServerError("database err: can't not convert to id", err)
		}
		MaDauSach, err := entity.StringToID(MaDauSachDB)
		if err != nil {
			return nil, coreerror.NewInternalServerError("database err: can't not convert to id", err)
		}

		dauSach, err := repo.dauSachRepository.GetDauSach(MaDauSach)

		if err != nil {
			return nil, err
		}

		sach := &entity.Sach{
			MaSach:     MaSach,
			DauSach:    dauSach,
			NhaXuatBan: NhaXuatBan,
			TriGia:     TriGia,
			NamXuatBan: NamXuatBan,
			TinhTrang:  TinhTrang,
		}
		ctPhieuNhap := &entity.CtPhieuNhap{
			Sach:   sach,
			DonGia: DonGia,
		}
		danhSachCtPhieuNhap = append(danhSachCtPhieuNhap, ctPhieuNhap)
	}
	return danhSachCtPhieuNhap, nil
}

func (repo *PhieuNhapRepository) GetPhieuNhap(maPhieuNhap *entity.ID) (_ *entity.PhieuNhap, err error) {
	tx := repo.db.MustBegin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// Query phieu nhap

	stmt, err := tx.Prepare(`SELECT MaPhieuNhap, NgayNhap, TongTien FROM PhieuNhap WHERE MaPhieuNhap = ?`)
	if err != nil {
		return nil, coreerror.NewInternalServerError("database err: can't not prepare query", err)
	}
	row := stmt.QueryRow(maPhieuNhap.String())
	var phieuNhap = &entity.PhieuNhap{}
	s := reflect.ValueOf(phieuNhap).Elem()
	numCols := s.NumField()
	columns := make([]interface{}, numCols)
	for i := 1; i < numCols; i++ {
		field := s.Field(i)
		columns[i] = field.Addr().Interface()
	}
	columns[0] = utils.Ptr("")
	err = row.Scan(columns[0 : len(columns)-1]...)
	if err == sql.ErrNoRows {
		return nil, coreerror.NewNotFoundError("phieu nhap not found", err)
	} else if err != nil {
		return nil, coreerror.NewInternalServerError("database err: can't not scan row", err)
	}
	phieuNhap.MaPhieuNhap = maPhieuNhap

	phieuNhap.CtPhieuNhap, err = repo.getCtPhieuNhapByPhieuNhap(phieuNhap.MaPhieuNhap.String())
	if err != nil {
		return nil, err
	}
	return phieuNhap, nil
}

func (repo *PhieuNhapRepository) CreatePhieuNhap(phieuNhap *entity.PhieuNhap) (_ *entity.PhieuNhap, err error) {
	tx := repo.db.MustBegin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	phieuNhapExec := `INSERT INTO PhieuNhap(MaPhieuNhap, NgayNhap, TongTien) VALUES (?, ?, ?)`
	_, err = tx.Exec(phieuNhapExec, phieuNhap.MaPhieuNhap.String(), phieuNhap.NgayLap, phieuNhap.TongTien)
	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not create phieu nhap", err)
	}

	return phieuNhap, nil
}

func (repo *PhieuNhapRepository) UpdatePhieuNhap(phieuNhap *entity.PhieuNhap) (_ *entity.PhieuNhap, err error) {
	tx := repo.db.MustBegin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	phieuNhapExec := `UPDATE PhieuNhap SET NgayNhap = ? WHERE MaPhieuNhap = ?`
	_, err = tx.Exec(phieuNhapExec, phieuNhap.NgayLap, phieuNhap.MaPhieuNhap.String())
	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not create phieu nhap", err)
	}

	return phieuNhap, nil
}

func (repo *PhieuNhapRepository) RemovePhieuNhap(phieuNhap *entity.PhieuNhap) (err error) {
	tx := repo.db.MustBegin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	phieuNhapExec := `DELETE FROM PhieuNhap WHERE MaPhieuNhap = ?`
	_, err = tx.Exec(phieuNhapExec, phieuNhap.MaPhieuNhap.String())
	if err != nil {
		if driverError, ok := err.(*mysql.MySQLError); ok {
			return DriverErrorHandling(driverError)
		}
		return coreerror.NewInternalServerError("database error: can't not delete phieu nhap", err)
	}
	return nil
}

func (repo *PhieuNhapRepository) AddChiTietPhieuNhap(maPhieuNhap *entity.ID, ctPhieuNhap *entity.CtPhieuNhap) (_ *entity.CtPhieuNhap, err error) {
	tx := repo.db.MustBegin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	sachExec := `INSERT INTO Sach (MaSach, MaDauSach, NhaXuatBan, TriGia, NamXuatBan, TinhTrang) VALUES (?, ?, ?, ? ,?, ?);`

	_, err = tx.Exec(sachExec, ctPhieuNhap.Sach.MaSach.String(), ctPhieuNhap.Sach.DauSach.MaDauSach.String(), ctPhieuNhap.Sach.NhaXuatBan, ctPhieuNhap.Sach.TriGia, ctPhieuNhap.Sach.NamXuatBan, ctPhieuNhap.Sach.TinhTrang)
	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not create sach", err)
	}
	ctPhieuNhapExec := `INSERT INTO Ct_PhieuNhap(MaPhieuNhap, MaSach, DonGia) VALUES (?, ?, ?)`
	_, err = tx.Exec(ctPhieuNhapExec, maPhieuNhap.String(), ctPhieuNhap.Sach.MaSach.String(), ctPhieuNhap.DonGia)
	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not create ct phieu nhap", err)
	}
	return ctPhieuNhap, nil
}

func (repo *PhieuNhapRepository) RemoveChiTietPhieuNhap(maSach *entity.ID) (err error) {
	tx := repo.db.MustBegin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	removeCtPhieuNhapExec := `DELETE FROM Ct_PhieuNhap WHERE MaSach = ?`
	removeSachExec := `DELETE FROM Sach WHERE MaSach = ?`
	_, err = tx.Exec(removeCtPhieuNhapExec, maSach)
	if err != nil {
		if driverError, ok := err.(*mysql.MySQLError); ok {
			return DriverErrorHandling(driverError)
		}
		return coreerror.NewInternalServerError("database error: can't not remove ct Phieu nhap", err)
	}
	_, err = tx.Exec(removeSachExec, maSach)
	if err != nil {
		if driverError, ok := err.(*mysql.MySQLError); ok {
			return DriverErrorHandling(driverError)
		}
		return coreerror.NewInternalServerError("database error: can't not remove sach", err)
	}
	return nil
}
