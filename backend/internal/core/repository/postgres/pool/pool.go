package core_postgres_pool

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/0ScPro0/affiliate-system/internal/core/config"
)

type Pool interface{
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Close()
	OperationTimeout() time.Duration
}

type ConnectionPool struct {
	*pgxpool.Pool
	operationTimeout time.Duration
}

func NewConnectionPool(ctx context.Context, cfg config.Config) (*ConnectionPool, error){
	connectionString := cfg.Database.DBUrl

	pgxconfig, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return nil, fmt.Errorf("Parse pgxconfig: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, pgxconfig)
	if err != nil {
		return nil, fmt.Errorf("Create pgxpool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("Ping pgxpool: %w", err)
	}

	return &ConnectionPool{
		Pool: pool,
		operationTimeout: cfg.Database.Timeot,
	}, nil
}

func (p *ConnectionPool) OperationTimeout() time.Duration {
	return p.operationTimeout
}