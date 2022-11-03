package entity

import "time"

type ThuThu struct {
	MaThuThu    *ID
	Name        string
	NgaySinh    *time.Time
	Email       string
	PhoneNumber string
	Password    string
	Status      bool
	IsAdminRole bool
}

func NewThuThu(
	name string,
	ngaySinh *time.Time,
	email string,
	phoneNumber string,
	password string,
	status bool,
	isAdminRole bool,
) *ThuThu {
	return &ThuThu{
		Name:        name,
		NgaySinh:    ngaySinh,
		Email:       email,
		PhoneNumber: phoneNumber,
		Password:    password,
		Status:      status,
		IsAdminRole: isAdminRole,
	}
}

func (thuThu *ThuThu) IsValid() bool {
	return true
}
