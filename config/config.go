package config

import (
	"github.com/miladshalikar/cafe/repository/postgreSQL"
	usertokenauthservice "github.com/miladshalikar/cafe/service/user/token"
)

type ServerConfig struct {
	Port int `koanf:"port"`
}

type Config struct {
	Server   ServerConfig
	Postgres postgreSQL.Config
	Token    usertokenauthservice.Config
}
