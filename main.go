package main

import (
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := initRouter()
	router.Run(":8080")
}
