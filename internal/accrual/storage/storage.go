package storage

import (
	"context"
	"io"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Storage interface {
	io.Closer

	Accrual
}

// Accrual методы для CRUD в БД
type Accrual interface {
	CreateOrderWithGoods(ctx context.Context, order string, goods []*Goods) (uuid.UUID, error)
	CreateInvalidOrder(ctx context.Context, order string) error
	UpdateOrder(ctx context.Context, order *Order) error
	UpdateGoodAccrual(ctx context.Context, goodID uuid.UUID, accrual int) error
	CreateMatch(ctx context.Context, match *Match) (uuid.UUID, error)
	GetMatchByName(ctx context.Context, matchName string) (uuid.UUID, error)
	GetOrderWithGoodsByNumber(ctx context.Context, orderNumber string) (*OrderOut, error)
}

// OrderStatus тип статуса заказа
type OrderStatus uint8

const (
	Registered = iota
	Invalid
	Processing
	Processed
)

func (o OrderStatus) String() string {
	switch o {
	case Registered:
		return "registered"
	case Invalid:
		return "invalid"
	case Processing:
		return "processing"
	case Processed:
		return "processed"
	}
	return "unknow"
}

// Goods структура для создания записи в таблице goods
type Goods struct {
	MatchID uuid.UUID
	Price   int
}

// Order структура для обновления записи в таблице orders
type Order struct {
	OrderID uuid.UUID
	Status  OrderStatus
	Accrual int
}

// OrderOut структура для выгрузки записи из таблицы orders
type OrderOut struct {
	OrderNumber string
	Status      string
	Accrual     int
}

// Match структура для создания записи в таблице matches
type Match struct {
	MatchName string
	Reward    int
	Type      RewardType
}

type RewardType uint8

const (
	Percent = iota
	Natural
)

func (r RewardType) String() string {
	switch r {
	case Percent:
		return "percent"
	case Natural:
		return "natural"
	}
	return "unknow"
}
