package config

import (
	"time"

	"github.com/rs/zerolog"
)

type Config struct {
	Env     Env
	Port    int     `env:"PORT" envDefault:"8080"`
	Host    string  `env:"HOST" envDefault:""`
	Logger  Logger  `envPrefix:"LOG_"`
	DB      DB      `envPrefix:"DB_"`
	Http    Http    `envPrefix:"HTTP_"`
	Cors    Cors    `envPrefix:"CORS_"`
	Swagger Swagger `envPrefix:"SWAGGER_"`
}

type Logger struct {
	Level         zerolog.Level `env:"LEVEL" envDefault:"warn"`
	Path          string        `env:"PATH,required,notEmpty" envDefault:""`
	IsCompressed  bool          `env:"IS_COMPRESSED" envDefault:"true"`
	IsDailyRotate bool          `env:"IS_DAILY_ROTATE" envDefault:"true"`
	SleepDuration time.Duration `env:"SLEEP_DURATION" envDefault:"1s"`
}

type DB struct {
	Name               string        `env:"NAME,required,notEmpty"`
	Host               string        `env:"HOST" envDefault:"localhost"`
	Port               int           `env:"PORT" envDefault:"5432"`
	User               string        `env:"USER,required" envDefault:""`
	Password           string        `env:"PASSWORD,required" envDefault:""`
	ConnectionIdle     time.Duration `env:"CONNECTION_IDLE" envDefault:"1m"`
	ConnectionLifetime time.Duration `env:"CONNECTION_LIFETIME" envDefault:"5m"`
	MaxIdle            int           `env:"MAX_IDLE" envDefault:"30"`
	MaxOpen            int           `env:"MAX_OPEN" envDefault:"90"`
}

type Http struct {
	MaxUploadSize int `env:"MAX_UPLOAD_SIZE" envDefault:"10"` // IN MB
}

type Cors struct {
	AllowOrigins     string `env:"ALLOW_ORIGINS" envDefault:"*"`
	AllowMethods     string `env:"ALLOW_METHODS" envDefault:"GET,POST,PUT,DELETE,OPTIONS"`
	AllowHeaders     string `env:"ALLOW_HEADERS" envDefault:"Origin,Content-Type,Accept,Authorization"`
	AllowCredentials bool   `env:"ALLOW_CREDENTIALS" envDefault:"false"`
	ExposeHeaders    string `env:"EXPOSE_HEADERS" envDefault:"Content-Length,Content-Type,Authorization"`
}

type Swagger struct {
	Username string `env:"USERNAME" envDefault:""`
	Password string `env:"PASSWORD" envDefault:""`
}
