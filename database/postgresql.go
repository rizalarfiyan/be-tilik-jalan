package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"os"
	"sync"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/rizalarfiyan/be-tilik-jalan/config"
	"github.com/rizalarfiyan/be-tilik-jalan/logger"
	"github.com/rs/zerolog"
)

var oncePostgres sync.Once
var pgSql *sql.DB

func InitPostgresql(ctx context.Context, conf *config.Config) {
	oncePostgres.Do(func() {
		log := logger.Get("postgresql")

		log.Info().Msg("Auto migrating postgresql database")

		val := url.Values{}
		if conf.DB.SSL {
			val.Add("sslmode", "require")
		} else {
			val.Add("sslmode", "disable")
		}
		dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?%s", conf.DB.User, conf.DB.Password, conf.DB.Host, conf.DB.Port, conf.DB.Name, val.Encode())
		autoMigrate(log, dsn)

		db, err := sql.Open("pgx", dsn)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to connect to postgresql database")
		}

		db.SetConnMaxLifetime(conf.DB.ConnectionLifetime)
		db.SetMaxOpenConns(conf.DB.MaxOpen)
		db.SetMaxIdleConns(conf.DB.MaxIdle)
		db.SetConnMaxIdleTime(conf.DB.ConnectionIdle)

		err = db.PingContext(ctx)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to ping postgresql database")
		}

		log.Info().Msg("Postgresql database connected")

		pgSql = new(sql.DB)
		pgSql = db
	})
}

func autoMigrate(log *zerolog.Logger, dsn string) {
	baseDir := "database/migrations"
	files, err := os.ReadDir(baseDir)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Warn().Msg("Migration directory does not exist, skipping migration")
			return
		}

		log.Fatal().Err(err).Msg("Failed to read migration directory")
	}

	if len(files) == 0 {
		log.Info().Msg("No migration files found, skipping migration")
		return
	}

	m, err := migrate.New("file://"+baseDir, dsn)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create migration")
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal().Err(err).Msg("Failed to migrate")
	}
}

func GetPostgresql() *sql.DB {
	return pgSql
}
