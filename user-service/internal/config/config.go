package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct{
	Env string				`yaml:"env" env-default:"local"`
	StoragePath string		`yaml:"storage_path" env-required:"true"`
	GRPC GRPCConfig			`yaml:"grpc" env-required:"true"`
}

type GRPCConfig struct{
	Port int32				`yaml:"port" env-required:"true"`
	Timeout time.Duration	`yaml:"timeout" env-default:"1h"`
}


func MustLoad() *Config{
	configPath := fetchConfigPath()

	if _, err := os.Stat(configPath); os.IsNotExist(err){
		panic("config file does not exist: " + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil{
		panic("config path is empty: " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string{
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == ""{
		os.Getenv("CONFIG_PATH")
	}
	return res
}