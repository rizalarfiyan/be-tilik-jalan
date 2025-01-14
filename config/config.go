package config

import (
	"log"
	"sync"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

var conf *Config
var once sync.Once

func Init() {
	once.Do(func() {
		log.Println("Initiating configuration")
		environment := new(Env).FromEnv("ENV")
		isProduction := environment.IsProduction()
		isStaging := environment.IsStaging()

		if !(isProduction || isStaging) {
			log.Println("Loading .env file")
			if err := godotenv.Load(".env"); err != nil {
				log.Fatalf("Error loading .env file: %v", err)
			}
		}

		log.Println("Parsing environment variable")
		cfg, err := env.ParseAs[Config]()
		if err != nil {
			log.Fatalf("Failed to parse environment variable: %v", err)
		}

		cfg.Env = environment
		conf = &cfg
		log.Println("Configuration initiated")
	})
}

func Get() *Config {
	return conf
}
