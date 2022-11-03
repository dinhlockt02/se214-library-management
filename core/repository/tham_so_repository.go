package repository

type ThamSoRepository interface {
	GetThoiHanThe() uint
	GetTuoiToiThieu() uint
	GetTuoiToiDa() uint
	GetDefaultPassword() string
}
