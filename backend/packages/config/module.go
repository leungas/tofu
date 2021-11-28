package configuration

import "githubcom/leungas/tofu/providers/config/services"

type ConfigModule struct {
	Service *services.ConfigService
}

func New(service *services.ConfigService) ConfigModule {
	return ConfigModule{Service: service}
}
