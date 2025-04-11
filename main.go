package main

import (
	"fmt"
	"github.com/miladshalikar/cafe/config"
	"github.com/miladshalikar/cafe/repository/migrator"
)

func main() {

	fmt.Println("faasssaa")

	cfg := config.C()
	m := migrator.New(cfg.Postgres)
	m.Up()

}
