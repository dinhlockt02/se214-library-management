package mysql

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	coreerror "daijoubuteam.xyz/se214-library-management/core/error"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"reflect"
)

type LoaiDocGiaRepository struct {
	db *sqlx.DB
}

func NewLoaiDocGiaRepository(db *sqlx.DB) *LoaiDocGiaRepository {
	return &LoaiDocGiaRepository{
		db: db,
	}
}

func (repo *LoaiDocGiaRepository) CreateLoaiDocGia(loaiDocGia *entity.LoaiDocGia) (_ *entity.LoaiDocGia, err error) {
	tx := repo.db.MustBegin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	exec := `INSERT INTO LoaiDocGia (MaLoaiDocGia, TenLoaiDocGia, SoSachToiDaDuocMuon, ThoiGianMuonToiDa, TienPhatTheoNgay) VALUES (? , ?, ?, ?, ?);`
	_, err = tx.Exec(exec, loaiDocGia.MaLoaiDocGia.String(), loaiDocGia.TenLoaiDocGia, loaiDocGia.SoSachToiDaDuocMuon, loaiDocGia.ThoiGianMuonToiDa, loaiDocGia.TienPhatTheoNgay)
	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: insert new loai doc gia failed", err)
	}
	return loaiDocGia, nil
}

func (repo *LoaiDocGiaRepository) GetLoaiDocGia(maLoaiDocGia *entity.ID) (_ *entity.LoaiDocGia, err error) {
	tx := repo.db.MustBegin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	stmt, err := tx.Prepare("SELECT MaLoaiDocGia, TenLoaiDocGia, SoSachToiDaDuocMuon, TienPhatTheoNgay, ThoiGianMuonToiDa FROM LoaiDocGia WHERE MaLoaiDocGia = ?")
	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not prepare query", err)
	}
	row := stmt.QueryRow(maLoaiDocGia.String())

	maLoaiDocGiaDB := ""

	loaiDocGia := &entity.LoaiDocGia{}

	s := reflect.ValueOf(loaiDocGia).Elem()
	numsCol := s.NumField()
	columns := make([]interface{}, numsCol)
	for i, _ := range columns {
		field := s.Field(i)
		columns[i] = field.Addr().Interface()
	}

	columns[0] = &maLoaiDocGiaDB

	err = row.Scan(columns...)
	if err != nil {
		return nil, coreerror.NewNotFoundError("Loai doc gia not found", err)
	}

	loaiDocGia.MaLoaiDocGia, err = entity.StringToID(maLoaiDocGiaDB)
	if err != nil {
		return nil, coreerror.NewInternalServerError("databaser error: convert to id failed", err)
	}
	return loaiDocGia, nil
}

func (repo *LoaiDocGiaRepository) GetDanhSachLoaiDocGia() (_ []*entity.LoaiDocGia, err error) {
	tx := repo.db.MustBegin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	stmt, err := tx.Prepare("SELECT MaLoaiDocGia, TenLoaiDocGia, SoSachToiDaDuocMuon, TienPhatTheoNgay, ThoiGianMuonToiDa FROM LoaiDocGia")
	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not prepare query", err)
	}
	rows, err := stmt.Query()
	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not execute query", err)
	}
	defer rows.Close()
	danhSachLoaiDocGia := make([]*entity.LoaiDocGia, 0)
	for rows.Next() {
		maLoaiDocGia := ""
		loaiDocGia := &entity.LoaiDocGia{}
		s := reflect.ValueOf(loaiDocGia).Elem()
		numCols := s.NumField()
		columns := make([]interface{}, numCols)
		for index, _ := range columns {
			field := s.Field(index)
			columns[index] = field.Addr().Interface()
		}
		columns[0] = &maLoaiDocGia
		err := rows.Scan(columns...)
		if err != nil {
			return danhSachLoaiDocGia, coreerror.NewInternalServerError("database error: scan rows failed", err)
		}
		loaiDocGia.MaLoaiDocGia, err = entity.StringToID(maLoaiDocGia)
		if err != nil {
			fmt.Println(maLoaiDocGia)
			return danhSachLoaiDocGia, coreerror.NewInternalServerError("database error: scan rows failed", err)
		}
		danhSachLoaiDocGia = append(danhSachLoaiDocGia, loaiDocGia)
	}
	return danhSachLoaiDocGia, nil
}

func (repo *LoaiDocGiaRepository) UpdateLoaiDocGia(loaiDocGia *entity.LoaiDocGia) (_ *entity.LoaiDocGia, err error) {
	tx := repo.db.MustBegin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	exec := `UPDATE LoaiDocGia SET TenLoaiDocGia = ?, SoSachToiDaDuocMuon = ?, ThoiGianMuonToiDa = ?, TienPhatTheoNgay = ? WHERE MaLoaiDocGia = ?;`
	_, err = tx.Exec(exec, loaiDocGia.TenLoaiDocGia, loaiDocGia.SoSachToiDaDuocMuon, loaiDocGia.ThoiGianMuonToiDa, loaiDocGia.TienPhatTheoNgay, loaiDocGia.MaLoaiDocGia.String())
	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: insert new loai doc gia failed", err)
	}
	return loaiDocGia, nil
}

func (repo *LoaiDocGiaRepository) RemoveLoaiDocGia(maLoaiDocGia *entity.ID) (err error) {
	tx := repo.db.MustBegin()
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()
	exec := `DELETE FROM LoaiDocGia WHERE MaLoaiDocGia = ?;`
	_, err = tx.Exec(exec, maLoaiDocGia.String())
	if err != nil {
		if driverError, ok := err.(*mysql.MySQLError); ok {
			return DriverErrorHandling(driverError)
		}
		return coreerror.NewInternalServerError("database error: delete new loai doc gia failed", err)
	}
	return nil
}
