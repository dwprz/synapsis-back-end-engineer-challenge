package database

import (
	"book-service/src/common/log"
	"book-service/src/infrastructure/config"
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

func NewPostgres() *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(config.Conf.Postgres.Dsn)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "database.NewPostgres", "section": "pgxpool.ParseConfig"}).Fatal(err)
	}

	config.MaxConnLifetime = time.Duration(30 * time.Minute)
	config.MaxConnLifetimeJitter = time.Duration(5 * time.Minute)
	config.MaxConnIdleTime = time.Duration(15 * time.Minute)
	config.MaxConns = 100
	config.MinConns = 10
	config.HealthCheckPeriod = time.Duration(1 * time.Minute)
	config.ConnConfig.Tracer = &log.QueryTracer{}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "database.NewPostgres", "section": "pgxpool.NewWithConfig"}).Fatal(err)
	}

	return pool
}
