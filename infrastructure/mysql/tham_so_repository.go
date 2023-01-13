package mysql

import "github.com/jmoiron/sqlx"

type ThamSoRepository struct {
	db *sqlx.DB
}

func NewThamSoRepository(db *sqlx.DB) *ThamSoRepository {
	return &ThamSoRepository{
		db: db,
	}
}

func (r *ThamSoRepository) GetThoiHanThe() (uint, error) {
	return 6, nil
}
func (r *ThamSoRepository) GetTuoiToiThieu() (uint, error) {
	return 18, nil
}
func (r *ThamSoRepository) GetTuoiToiDa() (uint, error) {
	return 55, nil
}
func (r *ThamSoRepository) GetDefaultPassword() (string, error) {
	return "12345678", nil
}
