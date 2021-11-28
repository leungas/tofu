package config

import "github.com/leungas/tofu/packages/config/services"

type ConfigModule struct {
	Service *services.ConfigService
}

func New(service *services.ConfigService) ConfigModule {
	return ConfigModule{Service: service}
}
