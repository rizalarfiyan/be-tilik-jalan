package config

import (
	"errors"
	"log"
	"net/url"
	"reflect"
	"sync"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
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

		opts := env.Options{
			FuncMap: map[reflect.Type]env.ParserFunc{
				reflect.TypeOf(new(zerolog.Level)): func(v string) (interface{}, error) {
					level, err := zerolog.ParseLevel(v)
					if err != nil {
						return nil, err
					}
					return level, nil
				},
				reflect.TypeOf(new(url.URL)): func(v string) (interface{}, error) {
					u, err := url.Parse(v)
					if err != nil {
						return nil, err
					}
					return *u, nil
				},
			},
		}

		log.Println("Parsing environment variable")
		cfg, err := env.ParseAsWithOptions[Config](opts)
		if err != nil {
			aggErr := env.AggregateError{}
			if ok := errors.As(err, &aggErr); ok {
				for _, er := range aggErr.Errors {
					log.Println(er)
				}
				log.Fatalln("Environment is not valid. Please check the configuration")
			}

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
