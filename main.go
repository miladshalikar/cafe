package main

import (
	"fmt"
	liaraobjectstorage "github.com/miladshalikar/cafe/adapter/objectstorage/liara"
	"github.com/miladshalikar/cafe/config"
	httpserver "github.com/miladshalikar/cafe/delivery/http_server"
	categoryhandler "github.com/miladshalikar/cafe/delivery/http_server/category"
	itemhandler "github.com/miladshalikar/cafe/delivery/http_server/item"
	mediahandler "github.com/miladshalikar/cafe/delivery/http_server/media"
	userauthhandler "github.com/miladshalikar/cafe/delivery/http_server/user/auth"
	userprofilehandler "github.com/miladshalikar/cafe/delivery/http_server/user/profile"
	"github.com/miladshalikar/cafe/repository/cache/redis"
	mediaredis "github.com/miladshalikar/cafe/repository/cache/redis/media"
	"github.com/miladshalikar/cafe/repository/migrator"
	"github.com/miladshalikar/cafe/repository/postgreSQL"
	aclpostgresql "github.com/miladshalikar/cafe/repository/postgreSQL/acl"
	categorypostgresql "github.com/miladshalikar/cafe/repository/postgreSQL/category"
	itempostgresql "github.com/miladshalikar/cafe/repository/postgreSQL/item"
	mediapostgresql "github.com/miladshalikar/cafe/repository/postgreSQL/media"
	userpostgresql "github.com/miladshalikar/cafe/repository/postgreSQL/user"
	aclservice "github.com/miladshalikar/cafe/service/acl"
	categoryservice "github.com/miladshalikar/cafe/service/categoty"
	itemservice "github.com/miladshalikar/cafe/service/item"
	mediaservice "github.com/miladshalikar/cafe/service/media"
	userauthservice "github.com/miladshalikar/cafe/service/user/authservice"
	userprofileservice "github.com/miladshalikar/cafe/service/user/profile"
	usertokenauthservice "github.com/miladshalikar/cafe/service/user/token"
	categoryvalidator "github.com/miladshalikar/cafe/validator/category"
	itemvalidator "github.com/miladshalikar/cafe/validator/item"
	mediavalidator "github.com/miladshalikar/cafe/validator/media"
	userauthvalidator "github.com/miladshalikar/cafe/validator/user/auth"
)

func main() {

	fmt.Println("hello world")
	cfg := config.C()
	m := migrator.New(cfg.Postgres)
	m.Up()
	postgresql := postgreSQL.New(cfg.Postgres)
	db := postgresql.Conn()

	cach := redis.New(cfg.Redis)
	cache := cach.Conn()

	objectStorage := liaraobjectstorage.New(cfg.ObjectStorage)

	userDB := userpostgresql.New(db)
	tknSvc := usertokenauthservice.New(cfg.Token)
	userVld := userauthvalidator.New(userDB)
	userSvc := userauthservice.New(userDB, tknSvc)
	userHandler := userauthhandler.New(userSvc, userVld)

	aclDB := aclpostgresql.New(db)
	acl := aclservice.New(aclDB)

	profileSvc := userprofileservice.New(userDB)
	ProfileHandler := userprofilehandler.New(profileSvc, tknSvc, cfg.Token)

	mediaDB := mediapostgresql.New(db)
	rCache := mediaredis.New(cache)
	mediaVld := mediavalidator.New(mediaDB)
	mediaSvc := mediaservice.New(objectStorage, mediaDB)
	mediaHandler := mediahandler.New(mediaSvc, mediaVld, tknSvc, cfg.Token, acl, cfg.ObjectStorage)

	categoryDB := categorypostgresql.New(db)
	categoryVld := categoryvalidator.New(categoryDB)
	categorySvc := categoryservice.New(categoryDB, mediaSvc, rCache)
	categoryHandler := categoryhandler.New(categorySvc, categoryVld, tknSvc, cfg.Token, acl)

	itemDB := itempostgresql.New(db)
	itemVld := itemvalidator.New(itemDB)
	itemSvc := itemservice.New(itemDB, mediaSvc, rCache)
	itemHandler := itemhandler.New(itemSvc, itemVld, tknSvc, cfg.Token, acl)

	echoServer := httpserver.New(cfg, userHandler, ProfileHandler, categoryHandler, mediaHandler, itemHandler)
	echoServer.Serve()
}
