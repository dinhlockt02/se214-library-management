package mysql

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	coreerror "daijoubuteam.xyz/se214-library-management/core/error"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type TacGiaRepository struct {
	db *sqlx.DB
}

func NewTacGiaRepository(db *sqlx.DB) *TacGiaRepository {
	return &TacGiaRepository{
		db: db,
	}
}

func (repo *TacGiaRepository) CreateTacGia(tacGia *entity.TacGia) (_ *entity.TacGia, err error) {
	tx := repo.db.MustBegin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	exec := `INSERT INTO TacGia(MaTacGia, TenTacGia) VALUES (?, ?);`
	_, err = tx.Exec(exec, tacGia.MaTacGia.String(), tacGia.TenTacGia)
	if err != nil {
		return nil, coreerror.NewInternalServerError("databaser error: can't not create tac gia", err)
	}
	return tacGia, nil
}

func (repo *TacGiaRepository) GetTacGia(maTacGia *entity.ID) (_ *entity.TacGia, err error) {
	tx := repo.db.MustBegin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	stmt, err := tx.Prepare(`SELECT TenTacGia FROM TacGia WHERE MaTacGia = ?`)
	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not prepare query", err)
	}
	row := stmt.QueryRow(maTacGia.String())
	tacGia := &entity.TacGia{}
	err = row.Scan(&(tacGia.TenTacGia))
	if err == sql.ErrNoRows {
		return nil, coreerror.NewNotFoundError("tac gia not found", err)
	} else if err != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not scan row", err)
	}
	tacGia.MaTacGia = maTacGia
	return tacGia, nil
}

func (repo *TacGiaRepository) GetDanhSachTacGia() (_ []*entity.TacGia, err error) {
	tx := repo.db.MustBegin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	stmt, err := tx.Prepare(`SELECT MaTacGia, TenTacGia FROM TacGia`)
	if err != nil {
		return nil, coreerror.NewInternalServerError("databaser error: can't not prepare query", err)
	}
	rows, err := stmt.Query()
	if err != nil {
		return nil, coreerror.NewInternalServerError("databaser error: can't not execute query", err)
	}
	danhSachTacGia := make([]*entity.TacGia, 0)
	for rows.Next() {
		tacGia := &entity.TacGia{}
		var maTacGia string = ""
		rows.Scan(&maTacGia, &(tacGia.TenTacGia))
		tacGia.MaTacGia, err = entity.StringToID(maTacGia)
		if err != nil {
			return nil, coreerror.NewInternalServerError("database error: can't not convert to id", err)
		}
		danhSachTacGia = append(danhSachTacGia, tacGia)
	}
	return danhSachTacGia, nil
}

func (repo *TacGiaRepository) UpdateTacGia(tacGia *entity.TacGia) (_ *entity.TacGia, err error) {
	tx := repo.db.MustBegin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	exec := `UPDATE TacGia SET TenTacGia = ? WHERE MaTacGia = ?;`
	_, err = tx.Exec(exec, tacGia.TenTacGia, tacGia.MaTacGia.String())
	if err != nil {
		return nil, coreerror.NewInternalServerError("databaser error: can't not update tac gia", err)
	}
	return tacGia, nil
}

func (repo *TacGiaRepository) RemoveTacGia(maTacGia *entity.ID) (err error) {
	tx := repo.db.MustBegin()
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	row := tx.QueryRow(`SELECT COUNT(*) FROM CT_TacGia WHERE MaTacGia = ?`, maTacGia.String())
	var count int = 0
	err = row.Scan(&count)
	if err != nil {
		return coreerror.NewInternalServerError("database error: cant query number of rows", err)
	}
	if count > 0 {
		return coreerror.NewConflictError("can't not delete tac gia because dau sach has this tac gia", nil)
	}

	exec := `DELETE FROM TacGia WHERE MaTacGia = ?`
	_, err = tx.Exec(exec, maTacGia.String())

	if err != nil {
		if driverError, ok := err.(*mysql.MySQLError); ok {
			return DriverErrorHandling(driverError)
		}
		return coreerror.NewInternalServerError("databaser error: can't not delete tac gia", err)
	}
	return nil
}
