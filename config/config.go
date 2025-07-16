package config

import (
	liaraobjectstorage "github.com/miladshalikar/cafe/adapter/objectstorage/liara"
	"github.com/miladshalikar/cafe/repository/postgreSQL"
	usertokenauthservice "github.com/miladshalikar/cafe/service/user/token"
)

type ServerConfig struct {
	Port int `koanf:"port"`
}

type Config struct {
	Server        ServerConfig                `koanf:"server"`
	Postgres      postgreSQL.Config           `koanf:"postgres"`
	Token         usertokenauthservice.Config `koanf:"token"`
	ObjectStorage liaraobjectstorage.Config   `koanf:"object_storage"`
}
