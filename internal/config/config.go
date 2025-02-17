package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/wensiet/logmod"
	"log/slog"
	"os"
	"sync"
)

type Config struct {
	Env      string `yaml:"env" env-required:"true"`
	Database struct {
		Host string `yaml:"host" env-required:"true"`
		Port string `yaml:"port" env-required:"true"`
		User string `yaml:"user" env-required:"true"`
		Pass string `yaml:"pass" env-required:"true"`
		Name string `yaml:"name" env-required:"true"`
	} `yaml:"database"`
	Loki struct {
		Host string `yaml:"host" env-required:"true"`
		Port int    `yaml:"port" env-required:"true"`
	} `yaml:"loki"`
	Services struct {
		Products struct {
			Host string `yaml:"host" env-required:"true"`
			Port string `yaml:"port" env-required:"true"`
		} `yaml:"products"`
	} `yaml:"services"`
	GRPC struct {
		Port int `yaml:"port" env-required:"true"`
	}
}

var once sync.Once
var config Config

func mustLoad() {
	configPath := ""
	if os.Getenv("DOCKER_ENV") == "true" {
		configPath = "./config/config_docker.yaml"
	}
	if os.Getenv("TEST_ENV") == "true" {
		configPath = "./config/config_test.yaml"
	}
	if configPath == "" {
		configPath = "./config/config_local.yaml"
	}
	err := cleanenv.ReadConfig(configPath, &config)
	if err != nil {
		panic(err)
	}
}

func Get() Config {
	once.Do(mustLoad)
	return config
}

func GetLogger() *slog.Logger {
	conf := Get()
	return logmod.New(logmod.Options{
		Env:     conf.Env,
		Service: "eshop-orders-ms",
		Loki: struct {
			Host string
			Port int
		}{
			Host: conf.Loki.Host,
			Port: conf.Loki.Port,
		},
	})
}
