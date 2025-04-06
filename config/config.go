package config

import "cafe/repository/postgreSQL"

type ServerConfig struct {
	Port int `koanf:"port"`
}

type Config struct {
	Server   ServerConfig
	Postgres postgreSQL.Config
}
