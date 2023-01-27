package presenter

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	"time"
)

type ThuThu struct {
	MaThuThu    string     `json:"maThuThu"`
	HoTen       string     `json:"hoTen"`
	NgaySinh    *time.Time `json:"ngaySinh"`
	Email       string     `json:"email"`
	SoDienThoai string     `json:"soDienThoai"`
}

func NewThuThuPresenter(thuThu *entity.ThuThu) ThuThu {
	return ThuThu{
		MaThuThu:    thuThu.MaThuThu.String(),
		HoTen:       thuThu.Name,
		NgaySinh:    thuThu.NgaySinh,
		Email:       thuThu.Email,
		SoDienThoai: thuThu.PhoneNumber,
	}
}
