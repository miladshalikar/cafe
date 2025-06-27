package config

import "time"

const (
	defaultPrefix       = "MA_"
	defaultDelimiter    = "."
	defaultSeparator    = "__"
	defaultYamlFilePath = "config.yml"

	AuthMiddlewareContextKey   = "claims"
	AccessTokenSubject         = "ac"
	RefreshTokenSubject        = "rt"
	AccessTokenExpireDuration  = time.Hour * 24 * 7
	RefreshTokenExpireDuration = time.Hour * 24 * 15
)
