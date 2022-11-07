package thuthucommandlogin

import (
	"daijoubuteam.xyz/se214-library-management/wireimpl"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func Login(db *sqlx.DB, email string, password string) {

	authSerivce := wireimpl.InitAuthUsecase(db)

	token, err := authSerivce.Login(email, password)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("User token is: %v\n", *token)
}
