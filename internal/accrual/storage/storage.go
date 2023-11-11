package storage

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/sergeizaitcev/gophermart/deployments/accrual/migrations"
	"github.com/sergeizaitcev/gophermart/internal/accrual/config"
)

// Accrual методы для CRUD в БД
type Accrual interface {
	CreateOrder(ctx context.Context, order string) (uuid.UUID, error)
}

type Storage struct {
	pool *pgxpool.Pool
}

func NewStorage(ctx context.Context, c *config.Config) (*Storage, error) {
	err := migrationUp(ctx, c.DatabaseURI)
	if err != nil {
		return nil, fmt.Errorf("migration: %w", err)
	}

	pool, err := newPool(ctx, c.DatabaseURI)
	if err != nil {
		return nil, fmt.Errorf("connect to the database: %w", err)
	}

	s := &Storage{
		pool: pool,
	}

	return s, nil
}

func (s *Storage) Close() error {
	s.pool.Close()
	return nil
}

func (s *Storage) Accrual() Accrual {
	return NewAccrualPostgres(s.pool)
}

func newPool(ctx context.Context, dsn string) (pool *pgxpool.Pool, err error) {
	pool, err = pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("create a new connection: %w", err)
	}
	defer func() {
		if err != nil {
			pool.Close()
		}
	}()

	pingCtx, pingCancel := context.WithTimeout(ctx, 3*time.Second)
	defer pingCancel()

	err = pool.Ping(pingCtx)
	if err != nil {
		return nil, fmt.Errorf("database ping %w", err)
	}

	return pool, nil
}

// migrationUp по dsn определяет *sql.DB и запускает миграцию
func migrationUp(ctx context.Context, dsn string) error {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return fmt.Errorf("create a new connection: %w", err)
	}
	defer db.Close()

	err = migrations.Up(ctx, db)
	if err != nil {
		return fmt.Errorf("migration up: %w", err)
	}

	return nil
}