package main

import (
	"fmt"
	"github.com/miladshalikar/cafe/config"
	httpserver "github.com/miladshalikar/cafe/delivery/http_server"
	userauthhandler "github.com/miladshalikar/cafe/delivery/http_server/user/auth"
	userprofilehandler "github.com/miladshalikar/cafe/delivery/http_server/user/profile"
	"github.com/miladshalikar/cafe/repository/migrator"
	"github.com/miladshalikar/cafe/repository/postgreSQL"
	userpostgresql "github.com/miladshalikar/cafe/repository/postgreSQL/user"
	userauthservice "github.com/miladshalikar/cafe/service/user/authservice"
	userprofileservice "github.com/miladshalikar/cafe/service/user/profile"
	usertokenauthservice "github.com/miladshalikar/cafe/service/user/token"
	userauthvalidator "github.com/miladshalikar/cafe/validator/user/auth"
)

func main() {

	fmt.Println("hello worldd")
	cfg := config.C()
	m := migrator.New(cfg.Postgres)
	m.Up()
	pd := postgreSQL.New(cfg.Postgres)
	db := pd.Conn()

	repo := userpostgresql.New(db)
	tok := usertokenauthservice.New(cfg.Token)
	val := userauthvalidator.New(repo)
	ser := userauthservice.New(repo, tok)
	hand := userauthhandler.New(ser, val)
	aaa := userprofileservice.New(repo)
	handy := userprofilehandler.New(aaa, tok, cfg.Token)
	echoServer := httpserver.New(cfg, hand, handy)
	echoServer.Serve()
}
