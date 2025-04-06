package main

import (
	"cafe/config"
	"cafe/repository/postgreSQL"
	"fmt"
)

func main() {

	fmt.Println("fsss")

	cfg := config.C()
	postgreSQL.InitDb(cfg.Postgres)

	//api.InitServer(cfg.Server)
}
