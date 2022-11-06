package thuthucommandls

import (
	"fmt"
	"strings"

	"daijoubuteam.xyz/se214-library-management/infrastructure/mysql"
	"daijoubuteam.xyz/se214-library-management/infrastructure/service"
	thuthu "daijoubuteam.xyz/se214-library-management/usecase/thu_thu"
	"github.com/fatih/color"
	"github.com/jmoiron/sqlx"
	"github.com/rodaine/table"
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
	table.DefaultHeaderFormatter = func(format string, vals ...interface{}) string {
		return strings.ToUpper(fmt.Sprintf(format, vals...))
	}
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("Ma Thu Thu", "Name", "Ngay Sinh", "Email", "Phone")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt).WithPadding(5)
	for _, tt := range danhSachThuThu {
		tbl.AddRow(tt.MaThuThu, tt.Name, tt.NgaySinh, tt.Email, tt.PhoneNumber)
	}
	tbl.Print()
}
