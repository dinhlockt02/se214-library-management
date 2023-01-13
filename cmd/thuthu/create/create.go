package create

import (
	"daijoubuteam.xyz/se214-library-management/infrastructure/mysql"
	"daijoubuteam.xyz/se214-library-management/infrastructure/service"
	thuthu "daijoubuteam.xyz/se214-library-management/usecase/thu_thu"
	"daijoubuteam.xyz/se214-library-management/utils"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strconv"
	"strings"
	"time"
)

func CreateThuThu(db *sqlx.DB, name string, birth string, email string, PhoneNumber string, status bool, password string) {

	passwordHasher := service.NewBcryptPasswordHasher()
	thuThuRepository := mysql.NewThuThuRepository(db)
	thamSoRepository := mysql.NewThamSoRepository(db)

	thuThuService := thuthu.NewThuThuService(passwordHasher, thuThuRepository, thamSoRepository)

	_, err := thuThuService.CreateThuThu(name, utils.Ptr(StringToDate(birth)), email, PhoneNumber, status, true, password)
	if err != nil {
		fmt.Println("Create thu thu failed")
		return
	}
	fmt.Println("Create thu thu successful")
}

func StringToDate(date string) time.Time {
	splitedDate := strings.Split(date, "-")

	if len(splitedDate) != 3 {
		panic("Invalid date")
	}

	year, err := strconv.ParseUint(splitedDate[0], 10, 16)

	if err != nil {
		panic("Invalid date")
	}

	month, err := strconv.ParseInt(splitedDate[1], 10, 16)
	if err != nil {
		panic("Invalid date")
	}
	day, err := strconv.ParseInt(splitedDate[2], 10, 16)
	if err != nil {
		panic("Invalid date")
	}

	rsDate := time.Date(int(year), time.Month(month), int(day), 0, 0, 0, 0, time.UTC)

	return rsDate
}
