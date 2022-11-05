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
	panic("not implemented")
}
func (r *ThamSoRepository) GetTuoiToiThieu() (uint, error) {
	panic("not implemented")
}
func (r *ThamSoRepository) GetTuoiToiDa() (uint, error) {
	panic("not implemented")
}
func (r *ThamSoRepository) GetDefaultPassword() (string, error) {
	panic("not implemented")
}
