package logger

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/rizalarfiyan/be-tilik-jalan/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

var once sync.Once
var log zerolog.Logger

func Init(conf *config.Config) {
	once.Do(func() {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		zerolog.TimeFieldFormat = time.RFC3339

		baseDir, err := filepath.Abs("./")
		isValidBaseDir := err == nil

		var output io.Writer = zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: time.RFC3339,
			FormatMessage: func(i interface{}) string {
				return fmt.Sprintf("| %s |", i)
			},
			FormatCaller: func(i interface{}) string {
				fullPath := fmt.Sprintf("%s", i)
				if !strings.HasPrefix(fullPath, baseDir) || !isValidBaseDir {
					return path.Base(fullPath)
				}

				currentFile, err := filepath.Rel(baseDir, fullPath)
				if err != nil {
					return path.Base(fullPath)
				}

				return currentFile
			},
		}

		if (conf.Env.IsProduction() || conf.Env.IsStaging()) && conf.Logger.File {
			lumberjackLog := NewLumberjackLogger(conf)
			output = zerolog.MultiLevelWriter(output, lumberjackLog.Run())
		}

		log = zerolog.New(output).
			Level(conf.Logger.Level).
			With().
			Timestamp().
			Logger()
	})
}

func Get(types string) *zerolog.Logger {
	logs := log.With().Str("type", types).Caller().Logger()
	return &logs
}

func GetWithoutCaller(types string) *zerolog.Logger {
	logs := log.With().Str("type", types).Logger()
	return &logs
}
