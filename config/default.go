package config

import usertokenauthservice "github.com/miladshalikar/cafe/service/user/token"

func Default() Config {
	cfx := Config{
		Token: usertokenauthservice.Config{
			AccessExpirationTime:  AccessTokenExpireDuration,
			RefreshExpirationTime: RefreshTokenExpireDuration,
			AccessSubject:         AccessTokenSubject,
			RefreshSubject:        RefreshTokenSubject,
		},
	}

	return cfx
}
