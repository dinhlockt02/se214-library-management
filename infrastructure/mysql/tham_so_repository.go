package mysql

import (
	coreerror "daijoubuteam.xyz/se214-library-management/core/error"
	"github.com/jmoiron/sqlx"
)

type ThamSoRepository struct {
	db *sqlx.DB
}

func (t ThamSoRepository) GetThoiHanThe() (uint, error) {
	var thoiHanThe uint
	if err := t.db.QueryRowx(`SELECT ThoiHanThe FROM ThamSo`).Scan(&thoiHanThe); err != nil {
		return 0, coreerror.NewInternalServerError("database error: get thoi han the failed", err)
	}
	return thoiHanThe, nil
}

func (t ThamSoRepository) GetTuoiToiThieu() (uint, error) {
	var tuoiToiThieu uint
	if err := t.db.QueryRowx(`SELECT TuoiToiThieu FROM ThamSo`).Scan(&tuoiToiThieu); err != nil {
		return 0, coreerror.NewInternalServerError("database error: get tuoi toi thieu failed", err)
	}
	return tuoiToiThieu, nil
}

func (t ThamSoRepository) GetTuoiToiDa() (uint, error) {
	var tuoiToiDa uint
	if err := t.db.QueryRowx(`SELECT TuoiToiDa FROM ThamSo`).Scan(&tuoiToiDa); err != nil {
		return 0, coreerror.NewInternalServerError("database error: get tuoi toi da failed", err)
	}
	return tuoiToiDa, nil
}

func (t ThamSoRepository) GetDefaultPassword() (string, error) {
	//TODO implement me
	panic("implement me")
}

func (t ThamSoRepository) GetTienPhatTheoNgay() (uint, error) {
	//TODO implement me
	panic("implement me")
}

func (t ThamSoRepository) GetThoiGianMuonToiDa(u uint, err error) {
	//TODO implement me
	panic("implement me")
}

func NewThamSoRepository(db *sqlx.DB) *ThamSoRepository {
	return &ThamSoRepository{
		db: db,
	}
}
