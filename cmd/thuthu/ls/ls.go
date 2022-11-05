package thuthucommandls

import (
	"fmt"

	"daijoubuteam.xyz/se214-library-management/infrastructure/mysql"
	"daijoubuteam.xyz/se214-library-management/infrastructure/service"
	thuthu "daijoubuteam.xyz/se214-library-management/usecase/thu_thu"
	"github.com/jmoiron/sqlx"
)

func ListThuThu(db *sqlx.DB, email string, phone string) {
	thuThuService := thuthu.NewThuThuService(
		service.NewBcryptPasswordHasher(),
		mysql.NewThuThuRepository(db),
		mysql.NewThamSoRepository(db),
	)

	danhSachThuThu, err := thuThuService.GetDanhSachThuThu(email, phone)
	if err != nil {
		fmt.Println("error: ls thu thu failed")
		return
	}
	fmt.Printf("Ma Thu Thu\t\tName\t\tNgaySinh\t\tEmail\t\tPhone\n")
	for _, tt := range danhSachThuThu {
		fmt.Printf("%v\t\t%v\t\t%v\t\t%v\t\t%v\n", tt.MaThuThu, tt.Name, tt.NgaySinh, tt.Email, tt.PhoneNumber)
	}
}
