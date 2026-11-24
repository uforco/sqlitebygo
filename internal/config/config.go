package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Address string `yaml:"address" env:"ADDRESS" env-required:"true" env-default:"localhost:3000"`
}

// env-default:"production"
type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	StoragePath string `yaml:"storage_path" env:"STORAGE_PATH" env-required:"true"`
	HttpServer  `yaml:"http_server"`
}

func MustLoad(path ...string) *Config {
	var configPath string
	configPath = os.Getenv(".env")
	if len(path) > 0 {
		configPath = path[0]
	}

	if configPath == "" {

		flags := flag.String("config", "", "path to config file")
		flag.Parse()

		configPath = *flags

		if configPath == "" {
			log.Fatal("Config path is not found")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exist: %s", configPath)
	}

	// clean env
	var clg Config
	err := cleanenv.ReadConfig(configPath, &clg)
	if err != nil {
		log.Fatalf("Cannot read config: %s", err.Error())
	}

	return &clg
}
