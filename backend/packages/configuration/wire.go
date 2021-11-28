//go:build wireinject
// +build wireinject

package configuration

import (
	"githubcom/leungas/tofu/providers/config/services"

	"github.com/google/wire"
)

func Initialize(options []services.ConfigOptions) (ConfigModule, error) {
	wire.Build(services.New, New)
	return ConfigModule{}, nil
}
