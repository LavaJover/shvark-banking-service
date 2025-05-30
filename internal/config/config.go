package config

import (
	"log"
	"os"
	"github.com/ilyakaznacheev/cleanenv"
)

type BankingConfig struct {
	Env string 	`yaml:"env"`
	GRPCServer 	`yaml:"grpc_server"`
	BankingDB 	`yaml:"banking_db"`
	LogConfig 	`yaml:"log_config"`
}

type GRPCServer struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type BankingDB struct {
	Dsn string `yaml:"dsn"`
}

type LogConfig struct {
	LogLevel 	string 	`yaml:"log_level"`
	LogFormat 	string 	`yaml:"log_format"`
	LogOutput 	string 	`yaml:"log_output"`
}

func MustLoad() *BankingConfig {

	// Processing env config variable and file
	configPath := os.Getenv("BANKING_CONFIG_PATH")

	if configPath == ""{
		log.Fatalf("BANKING_CONFIG_PATH was not found\n")
	}

	if _, err := os.Stat(configPath); err != nil{
		log.Fatalf("failed to find config file: %v\n", err)
	}

	// YAML to struct object
	var cfg BankingConfig
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil{
		log.Fatalf("failed to read config file: %v", err)
	}

	return &cfg
}