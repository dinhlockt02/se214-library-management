package mysql

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	coreerror "daijoubuteam.xyz/se214-library-management/core/error"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DauSachRepository struct {
	db *sqlx.DB
}

func NewDauSachRepository(db *sqlx.DB) *DauSachRepository {
	return &DauSachRepository{
		db: db,
	}
}

func (repo *DauSachRepository) GetDanhSachDauSach() (_ []*entity.DauSach, err error) {
	tx := repo.db.MustBegin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	stmt, err := tx.Prepare(`
	SELECT DauSach.MaDauSach , TenDauSach, TacGia.MaTacGia , TenTacGia, TheLoai.MaTheLoai  ,TenTheLoai 
	FROM DauSach, CT_TacGia, CT_TheLoai, TacGia, TheLoai
	WHERE DauSach.MaDauSach = CT_TacGia.MaDauSach AND CT_TacGia.MaTacGia = TacGia.MaTacGia AND DauSach.MaDauSach = CT_TheLoai.MaDauSach AND CT_TheLoai.MaTheLoai = TheLoai.MaTheLoai`)
	type DauSachDB struct {
		MaDauSach  string
		TenDauSach string
		MaTacGia   string
		TenTacGia  string
		MaTheLoai  string
		TenTheLoai string
	}
	rows, err := stmt.Query()
	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not execute query", err)
	}
	DauSachMap := make(map[string]*entity.DauSach)
	for rows.Next() {
		var dauSachDB DauSachDB
		rows.Scan(&(dauSachDB.MaDauSach), &(dauSachDB.TenDauSach), &(dauSachDB.MaTacGia), &(dauSachDB.TenTacGia), &(dauSachDB.MaTheLoai), &(dauSachDB.TenTheLoai))
		maTacGia, err := entity.StringToID(dauSachDB.MaTacGia)
		if err != nil {
			return nil, coreerror.NewInternalServerError("database error: can't not convert to id", err)
		}
		maTheLoai, err := entity.StringToID(dauSachDB.MaTheLoai)
		if err != nil {
			return nil, coreerror.NewInternalServerError("database error: can't not convert to id", err)
		}
		maDauSach, err := entity.StringToID(dauSachDB.MaDauSach)
		if err != nil {
			fmt.Println(dauSachDB.MaDauSach)
			return nil, coreerror.NewInternalServerError("database error: can't not convert to id", err)
		}
		if _, ok := DauSachMap[dauSachDB.MaDauSach]; ok {
			tacGia := &entity.TacGia{
				MaTacGia:  maTacGia,
				TenTacGia: dauSachDB.TenTacGia,
			}
			theLoai := &entity.TheLoai{
				MaTheLoai:  maTheLoai,
				TenTheLoai: dauSachDB.TenTheLoai,
			}
			DauSachMap[dauSachDB.MaDauSach].TacGia = append(DauSachMap[dauSachDB.MaDauSach].TacGia, tacGia)
			DauSachMap[dauSachDB.MaDauSach].TheLoai = append(DauSachMap[dauSachDB.MaDauSach].TheLoai, theLoai)
		} else {
			tacGia := []*entity.TacGia{
				&entity.TacGia{
					MaTacGia:  maTacGia,
					TenTacGia: dauSachDB.TenTacGia,
				},
			}
			theLoai := []*entity.TheLoai{
				&entity.TheLoai{
					MaTheLoai:  maTheLoai,
					TenTheLoai: dauSachDB.TenTheLoai,
				},
			}
			DauSachMap[dauSachDB.MaDauSach] = &entity.DauSach{
				MaDauSach:  maDauSach,
				TenDauSach: dauSachDB.TenDauSach,
				TacGia:     tacGia,
				TheLoai:    theLoai,
			}
		}
	}

	danhSachDauSach := make([]*entity.DauSach, 0, len(DauSachMap))

	for _, values := range DauSachMap {
		uniqueTheLoai := make(map[string]*entity.TheLoai)
		for _, theLoai := range values.TheLoai {
			if _, ok := uniqueTheLoai[theLoai.MaTheLoai.String()]; ok {
				continue
			} else {
				uniqueTheLoai[theLoai.MaTheLoai.String()] = theLoai
			}
		}
		theLoai := make([]*entity.TheLoai, 0, len(uniqueTheLoai))
		for _, tl := range uniqueTheLoai {
			theLoai = append(theLoai, tl)
		}
		values.TheLoai = theLoai

		uniqueTacGia := make(map[string]*entity.TacGia)
		for _, tacGia := range values.TacGia {
			if _, ok := uniqueTheLoai[tacGia.MaTacGia.String()]; ok {
				continue
			} else {
				uniqueTacGia[tacGia.MaTacGia.String()] = tacGia
			}
		}
		tacGia := make([]*entity.TacGia, 0, len(uniqueTheLoai))
		for _, tg := range uniqueTacGia {
			tacGia = append(tacGia, tg)
		}
		values.TheLoai = theLoai
		values.TacGia = tacGia
		danhSachDauSach = append(danhSachDauSach, values)
	}
	return danhSachDauSach, nil
}

func (repo *DauSachRepository) GetDauSach(maDauSach *entity.ID) (_ *entity.DauSach, err error) {
	danhSachDauSach, err := repo.GetDanhSachDauSach()
	for _, dauSach := range danhSachDauSach {
		if dauSach.MaDauSach.String() == maDauSach.String() {
			return dauSach, nil
		}
	}
	return nil, coreerror.NewNotFoundError("dau sach not found", nil)
}

func (repo *DauSachRepository) CreateDauSach(dauSach *entity.DauSach) (_ *entity.DauSach, err error) {
	tx := repo.db.MustBegin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	exec := `INSERT INTO DauSach (MaDauSach, TenDauSach) VALUES (?, ?)`
	_, err = tx.Exec(exec, dauSach.MaDauSach.String(), dauSach.TenDauSach)
	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not insert dau sach", err)
	}
	for _, theLoai := range dauSach.TheLoai {
		exec = `INSERT INTO CT_TheLoai (MaDauSach, MaTheLoai) VALUES (?, ?)`
		_, err = tx.Exec(exec, dauSach.MaDauSach.String(), theLoai.MaTheLoai.String())
		if err != nil {
			return nil, coreerror.NewInternalServerError("database error: can't not insert Ct_TheLoai", err)
		}
	}
	for _, tacGia := range dauSach.TacGia {
		exec = `INSERT INTO CT_TacGia (MaDauSach, MaTacGia) VALUES (?, ?)`
		_, err = tx.Exec(exec, dauSach.MaDauSach.String(), tacGia.MaTacGia.String())
		if err != nil {
			return nil, coreerror.NewInternalServerError("database error: can't not inset Ct_TacGia", err)
		}
	}
	return dauSach, nil
}

func (repo *DauSachRepository) UpdateDauSach(dauSach *entity.DauSach) (_ *entity.DauSach, err error) {
	tx := repo.db.MustBegin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	exec := `UPDATE DauSach SET TenDauSach = ? WHERE MaDauSach = ?`
	_, err = tx.Exec(exec, dauSach.TenDauSach, dauSach.MaDauSach.String())
	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not update dau sach", err)
	}
	exec = `DELETE FROM CT_TheLoai WHERE MaDauSach = ?`
	_, err = tx.Exec(exec, dauSach.MaDauSach.String())
	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not remove Ct_TheLoai", err)
	}
	for _, theLoai := range dauSach.TheLoai {
		exec = `INSERT INTO CT_TheLoai (MaDauSach, MaTheLoai) VALUES (?, ?)`
		_, err = tx.Exec(exec, dauSach.MaDauSach.String(), theLoai.MaTheLoai.String())
		if err != nil {
			return nil, coreerror.NewInternalServerError("database error: can't not insert Ct_TheLoai", err)
		}
	}
	exec = `DELETE FROM CT_TacGia WHERE MaDauSach = ?`
	_, err = tx.Exec(exec, dauSach.MaDauSach.String())
	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not remove Ct_TacGia", err)
	}
	for _, tacGia := range dauSach.TacGia {
		exec = `INSERT INTO CT_TacGia (MaDauSach, MaTacGia) VALUES (?, ?)`
		_, err = tx.Exec(exec, dauSach.MaDauSach.String(), tacGia.MaTacGia.String())
		if err != nil {
			return nil, coreerror.NewInternalServerError("database error: can't not insert Ct_TacGia", err)
		}
	}
	return dauSach, nil
}

func (repo *DauSachRepository) RemoveDauSach(maDauSach *entity.ID) (err error) {
	tx := repo.db.MustBegin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	exec := `DELETE FROM CT_TheLoai WHERE MaDauSach = ?`
	_, err = tx.Exec(exec, maDauSach.String())
	if err != nil {
		if driverError, ok := err.(*mysql.MySQLError); ok {
			return DriverErrorHandling(driverError)
		}
		return coreerror.NewInternalServerError("database error: can't not delete Ct_TheLoai", err)
	}
	exec = `DELETE FROM CT_TacGia WHERE MaDauSach = ?`
	_, err = tx.Exec(exec, maDauSach.String())
	if err != nil {
		if driverError, ok := err.(*mysql.MySQLError); ok {
			return DriverErrorHandling(driverError)
		}
		return coreerror.NewInternalServerError("database error: can't not delete Ct_TheLoai", err)
	}
	exec = `DELETE FROM DauSach WHERE MaDauSach = ?`
	_, err = tx.Exec(exec, maDauSach.String())
	if err != nil {
		if driverError, ok := err.(*mysql.MySQLError); ok {
			return DriverErrorHandling(driverError)
		}
		return coreerror.NewInternalServerError("database error: can't not delete Ct_TheLoai", err)
	}
	return nil
}
