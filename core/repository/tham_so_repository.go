package repository

type ThamSoRepository interface {
	GetThoiHanThe() (uint, error)
	GetTuoiToiThieu() (uint, error)
	GetTuoiToiDa() (uint, error)
	GetDefaultPassword() (string, error)
}
