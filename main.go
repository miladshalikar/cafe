package main

import (
	"fmt"
	"github.com/miladshalikar/cafe/config"
	httpserver "github.com/miladshalikar/cafe/delivery/http_server"
	categoryhandler "github.com/miladshalikar/cafe/delivery/http_server/category"
	userauthhandler "github.com/miladshalikar/cafe/delivery/http_server/user/auth"
	userprofilehandler "github.com/miladshalikar/cafe/delivery/http_server/user/profile"
	"github.com/miladshalikar/cafe/repository/migrator"
	"github.com/miladshalikar/cafe/repository/postgreSQL"
	aclpostgresql "github.com/miladshalikar/cafe/repository/postgreSQL/acl"
	categorypostgresql "github.com/miladshalikar/cafe/repository/postgreSQL/category"
	userpostgresql "github.com/miladshalikar/cafe/repository/postgreSQL/user"
	aclservice "github.com/miladshalikar/cafe/service/acl"
	categoryservice "github.com/miladshalikar/cafe/service/categoty"
	userauthservice "github.com/miladshalikar/cafe/service/user/authservice"
	userprofileservice "github.com/miladshalikar/cafe/service/user/profile"
	usertokenauthservice "github.com/miladshalikar/cafe/service/user/token"
	categoryvalidator "github.com/miladshalikar/cafe/validator/category"
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

	rrr := categorypostgresql.New(db)
	nn := categoryvalidator.New(rrr)
	sss := categoryservice.New(rrr)
	aclr := aclpostgresql.New(db)
	acl := aclservice.New(aclr)
	hhh := categoryhandler.New(sss, nn, tok, cfg.Token, acl)

	echoServer := httpserver.New(cfg, hand, handy, hhh)
	echoServer.Serve()
}
