package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	Redis      redis      `yaml:"redis"`
	HttpServer httpServer `yaml:"http_server"`
	GrpcServer grpcServer `yaml:"grpc_server"`
	//	LoggerConfig  `yaml:"logger_config"`
}

type httpServer struct {
	Port      int    `yaml:"port"`
	ElkDomain string `yaml:"elk_domain"`
}

type grpcServer struct {
	Port int `yaml:"port"`
}

type redis struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

// InitConfig загружает конфигурацию из файла и обрабатывает переменные окружения
func InitConfig(filePath string) (*Config, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML config: %w", err)
	}

	//if err := processEnvVariables(&config); err != nil {
	//	return nil, err
	//}
	//
	//if err := validateConfig(&config); err != nil {
	//	return nil, err
	//}

	return &config, nil
}
