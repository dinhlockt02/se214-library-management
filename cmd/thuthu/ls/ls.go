package thuthucommandls

import (
	"fmt"
	"strings"

	"daijoubuteam.xyz/se214-library-management/wireimpl"
	"github.com/fatih/color"
	"github.com/jmoiron/sqlx"
	"github.com/rodaine/table"
)

func ListThuThu(db *sqlx.DB) {
	thuThuService := wireimpl.InitThuThuUsecase(db)

	danhSachThuThu, err := thuThuService.GetDanhSachThuThu()
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
