package config

import "github.com/miladshalikar/cafe/repository/postgreSQL"

type ServerConfig struct {
	Port int `koanf:"port"`
}

type Config struct {
	Server   ServerConfig
	Postgres postgreSQL.Config
}
