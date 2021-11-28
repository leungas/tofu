package services

import "github.com/spf13/viper"

type Configuration struct {
	Name string
	Data *viper.Viper
}

type ConfigOptions struct {
	Name string
	File string
	Path string
	Type string
}

type ConfigService struct {
	repository []Configuration
}

func (s *ConfigService) load(options []ConfigOptions) {
	for _, defs := range options {
		dataset := viper.New()
		dataset.SetConfigFile(defs.File)
		dataset.SetConfigType(defs.Type)
		dataset.AddConfigPath(defs.Path)
		if err := dataset.ReadInConfig(); err != nil {

		}
		s.repository = append(s.repository, Configuration{Name: defs.Name, Data: dataset})
	}
}

func New(options []ConfigOptions) *ConfigService {
	instance := ConfigService{repository: []Configuration{}}
	instance.load(options)
	return &instance
}

func (s *ConfigService) Get(target string, value string) interface{} {
	var data *viper.Viper

	for _, dataset := range s.repository {
		if dataset.Name == target {
			data = dataset.Data
		}
	}

	if data == nil {
		// throw exception
	}

	return data.Get(value)
}
