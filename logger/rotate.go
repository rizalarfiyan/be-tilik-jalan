package logger

import (
	"sync"
	"time"

	"github.com/rizalarfiyan/be-tilik-jalan/config"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LumberjackLog interface {
	Run() *lumberjack.Logger
}

type lumberjackLog struct {
	LogPath       string
	CompressLog   bool
	DailyRotate   bool
	SleepDuration time.Duration
	lastLogDate   string
	lumberjackLog *lumberjack.Logger
}

func NewLumberjackLogger(conf *config.Config) LumberjackLog {
	return &lumberjackLog{
		LogPath:       conf.Logger.Path,
		DailyRotate:   conf.Logger.IsDailyRotate,
		CompressLog:   conf.Logger.IsCompressed,
		SleepDuration: conf.Logger.SleepDuration,
	}
}

func (l *lumberjackLog) Run() *lumberjack.Logger {
	l.lumberjackLog = &lumberjack.Logger{
		Filename:  l.LogPath,
		Compress:  l.CompressLog,
		LocalTime: true,
	}

	l.lastLogDate = time.Now().Format(time.DateOnly)

	if l.DailyRotate {
		wg := &sync.WaitGroup{}

		wg.Add(1)
		go func() {
			l.rotate()
		}()

	}

	return l.lumberjackLog
}

func (l *lumberjackLog) rotate() {
	ticker := time.NewTicker(l.SleepDuration)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if l.lumberjackLog == nil {
				continue
			}

			now := time.Now().Format(time.DateOnly)
			if l.lastLogDate != now {
				log.Info().Msg("Rotating log file")
				l.lastLogDate = now
				err := l.lumberjackLog.Rotate()
				if err != nil {
					log.Error().Err(err).Msg("Failed to rotate log file")
				}
			}
		}
	}
}
