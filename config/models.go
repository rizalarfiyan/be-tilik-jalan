package config

import (
	"time"

	"github.com/rs/zerolog"
)

type Config struct {
	Env    Env
	Port   int    `env:"PORT" envDefault:"8080"`
	Host   string `env:"HOST" envDefault:""`
	Logger Logger `envPrefix:"LOG_"`
}

type Logger struct {
	Level         zerolog.Level `env:"LEVEL" envDefault:"warn"`
	Path          string        `env:"PATH,required,notEmpty" envDefault:""`
	IsCompressed  bool          `env:"IS_COMPRESSED" envDefault:"true"`
	IsDailyRotate bool          `env:"IS_DAILY_ROTATE" envDefault:"true"`
	SleepDuration time.Duration `env:"SLEEP_DURATION" envDefault:"1s"`
}
