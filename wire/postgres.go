package wire

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var PostgresConfigSingleton = NewSingleton(func() *pgxpool.Config {
	const defaultMaxConns = int32(4)
	const defaultMinConns = int32(0)
	const defaultMaxConnLifetime = time.Hour
	const defaultMaxConnIdleTime = time.Minute * 30
	const defaultHealthCheckPeriod = time.Minute
	const defaultConnectTimeout = time.Second * 5

	envSingleton := EnvSingleton.Get()

	dbConfig, err := pgxpool.ParseConfig(fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		envSingleton.POSTGRES_USER,
		envSingleton.POSTGRES_PASS,
		envSingleton.POSTGRES_HOST,
		envSingleton.POSTGRES_PORT,
		envSingleton.POSTGRES_NAME))
	if err != nil {
		panic(err)
	}

	dbConfig.MaxConns = defaultMaxConns
	dbConfig.MinConns = defaultMinConns
	dbConfig.MaxConnLifetime = defaultMaxConnLifetime
	dbConfig.MaxConnIdleTime = defaultMaxConnIdleTime
	dbConfig.HealthCheckPeriod = defaultHealthCheckPeriod
	dbConfig.ConnConfig.ConnectTimeout = defaultConnectTimeout

	return dbConfig
})

var PgxPoolSingleton = NewSingleton(func() *pgxpool.Pool {
	var err error
	pgxPool, err := pgxpool.NewWithConfig(
		context.Background(),
		PostgresConfigSingleton.Get(),
	)
	if err != nil {
		panic(err)
	}

	return pgxPool
})
