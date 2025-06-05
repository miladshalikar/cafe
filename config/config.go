package config

import (
	"github.com/miladshalikar/cafe/repository/postgreSQL"
	tokenservice "github.com/miladshalikar/cafe/service/token"
)

type ServerConfig struct {
	Port int `koanf:"port"`
}

type Config struct {
	Server   ServerConfig
	Postgres postgreSQL.Config
	Token    tokenservice.Config
}
