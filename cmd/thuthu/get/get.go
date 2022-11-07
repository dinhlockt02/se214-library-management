package thuthucommandget

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	"daijoubuteam.xyz/se214-library-management/wireimpl"
	"fmt"
	"github.com/fatih/color"
	"github.com/jmoiron/sqlx"
	"github.com/rodaine/table"
)

func GetThuThu(db *sqlx.DB, maThuThuStr string) {
	thuThuService := wireimpl.InitThuThuUsecase(db)

	maThuThu, err := entity.StringToID(maThuThuStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	thuThu, err := thuThuService.GetThuThu(maThuThu)
	if err != nil {
		fmt.Println(err)
		return
	}
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("Ma Thu Thu", "Name", "Ngay Sinh", "Email", "Phone")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt).WithPadding(5)
	tbl.AddRow(thuThu.MaThuThu, thuThu.Name, thuThu.NgaySinh, thuThu.Email, thuThu.PhoneNumber)
	tbl.Print()
}
