package admincreate

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"daijoubuteam.xyz/se214-library-management/infrastructure/mysql"
	"daijoubuteam.xyz/se214-library-management/infrastructure/service"
	thuthu "daijoubuteam.xyz/se214-library-management/usecase/thu_thu"
	"daijoubuteam.xyz/se214-library-management/utils"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
)

func CreateAdminCommand(db *sqlx.DB) *cobra.Command {

	var name string
	var birth string
	var email string
	var PhoneNumber string
	var status bool
	var password string

	command := &cobra.Command{
		Use:   `create`,
		Short: `Create new admin`,
		Run: func(cmd *cobra.Command, args []string) {
			CreateAdmin(db, name, birth, email, PhoneNumber, status, password)
		},
	}

	command.PersistentFlags().StringVarP(&name, "name", "n", "", "Admin's name")
	command.PersistentFlags().StringVarP(&birth, "birth", "b", "2022-12-31", "Admin's birthday")
	command.PersistentFlags().StringVar(&PhoneNumber, "phone", "", "Admin's phone number")
	command.PersistentFlags().BoolVar(&status, "enable", false, "Admin's enable status")
	command.PersistentFlags().StringVarP(&email, "email", "e", "", "Admin's email")
	command.MarkFlagRequired("email")
	command.PersistentFlags().StringVarP(&password, "password", "p", "", "Admin's password")
	command.MarkFlagRequired("password")
	return command
}

func CreateAdmin(db *sqlx.DB, name string, birth string, email string, PhoneNumber string, status bool, password string) {

	passwordHasher := service.NewBcryptPasswordHasher()
	thuThuRepository := mysql.NewThuThuRepository(db)
	thamSoRepository := mysql.NewThamSoRepository(db)

	thuThuService := thuthu.NewThuThuService(passwordHasher, thuThuRepository, thamSoRepository)

	_, err := thuThuService.CreateThuThu(name, utils.Ptr(StringToDate(birth)), email, PhoneNumber, status, true, password)
	if err != nil {
		fmt.Println("Create admin failed")
		return
	}
	fmt.Println("Create admin successful")
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
