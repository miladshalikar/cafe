package main

import (
	"fmt"
	liaraobjectstorage "github.com/miladshalikar/cafe/adapter/objectstorage/liara"
	"github.com/miladshalikar/cafe/config"
	httpserver "github.com/miladshalikar/cafe/delivery/http_server"
	categoryhandler "github.com/miladshalikar/cafe/delivery/http_server/category"
	mediahandler "github.com/miladshalikar/cafe/delivery/http_server/media"
	userauthhandler "github.com/miladshalikar/cafe/delivery/http_server/user/auth"
	userprofilehandler "github.com/miladshalikar/cafe/delivery/http_server/user/profile"
	"github.com/miladshalikar/cafe/repository/cache/redis"
	mediaredis "github.com/miladshalikar/cafe/repository/cache/redis/media"
	"github.com/miladshalikar/cafe/repository/migrator"
	"github.com/miladshalikar/cafe/repository/postgreSQL"
	aclpostgresql "github.com/miladshalikar/cafe/repository/postgreSQL/acl"
	categorypostgresql "github.com/miladshalikar/cafe/repository/postgreSQL/category"
	mediapostgresql "github.com/miladshalikar/cafe/repository/postgreSQL/media"
	userpostgresql "github.com/miladshalikar/cafe/repository/postgreSQL/user"
	aclservice "github.com/miladshalikar/cafe/service/acl"
	categoryservice "github.com/miladshalikar/cafe/service/categoty"
	mediaservice "github.com/miladshalikar/cafe/service/media"
	userauthservice "github.com/miladshalikar/cafe/service/user/authservice"
	userprofileservice "github.com/miladshalikar/cafe/service/user/profile"
	usertokenauthservice "github.com/miladshalikar/cafe/service/user/token"
	categoryvalidator "github.com/miladshalikar/cafe/validator/category"
	mediavalidator "github.com/miladshalikar/cafe/validator/media"
	userauthvalidator "github.com/miladshalikar/cafe/validator/user/auth"
)

func main() {

	fmt.Println("hello worldd")
	cfg := config.C()
	m := migrator.New(cfg.Postgres)
	m.Up()
	pd := postgreSQL.New(cfg.Postgres)
	db := pd.Conn()

	cach := redis.New(cfg.Redis)
	cache := cach.Conn()

	objectStorage := liaraobjectstorage.New(cfg.ObjectStorage)

	repo := userpostgresql.New(db)
	tknSvc := usertokenauthservice.New(cfg.Token)
	val := userauthvalidator.New(repo)
	ser := userauthservice.New(repo, tknSvc)
	hand := userauthhandler.New(ser, val)

	aclr := aclpostgresql.New(db)
	acl := aclservice.New(aclr)

	aaa := userprofileservice.New(repo)
	handy := userprofilehandler.New(aaa, tknSvc, cfg.Token)

	mediaDB := mediapostgresql.New(db)
	mediaVld := mediavalidator.New(mediaDB)
	mediaSvc := mediaservice.New(objectStorage, mediaDB)
	mediaHandler := mediahandler.New(mediaSvc, mediaVld, tknSvc, cfg.Token, acl)

	rrr := categorypostgresql.New(db)
	ccache := mediaredis.New(cache)
	nn := categoryvalidator.New(rrr)
	sss := categoryservice.New(rrr, mediaSvc, ccache)

	hhh := categoryhandler.New(sss, nn, tknSvc, cfg.Token, acl)

	echoServer := httpserver.New(cfg, hand, handy, hhh, mediaHandler)
	echoServer.Serve()
}
