package main

import (
	"fmt"
	"github.com/miladshalikar/cafe/config"
	"github.com/miladshalikar/cafe/repository/postgreSQL"
)

func main() {

	fmt.Println("fssssaa")

	cfg := config.C()
	postgreSQL.InitDb(cfg.Postgres)

	//api.InitServer(cfg.Server)
}
