package main

import (
	"fmt"

	"daijoubuteam.xyz/se214-library-management/config"
	"daijoubuteam.xyz/se214-library-management/utils"
)

func main() {

	db := utils.ConnectDB(config.DevConfig)

	fmt.Printf("Connect to database successful: %v \n", db)

	// repo := inmemory.InMemoryBookRepo{}
	// bookService := book.NewService(&repo)
	// r := gin.Default()
	// handler.MakeBookHandler(r, bookService)

	// r.Run(":8080")
}
