//go:build wireinject
// +build wireinject

package config

import (
	"github.com/google/wire"
	"github.com/leungas/tofu/packages/config/services"
)

func Initialize(options []services.ConfigOptions) (ConfigModule, error) {
	wire.Build(New, services.New)
	return ConfigModule{}, nil
}
