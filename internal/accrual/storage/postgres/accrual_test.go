package postgres

import (
	"context"
	"testing"

	"github.com/google/uuid"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/sergeizaitcev/gophermart/deployments/accrual/migrations"
	"github.com/sergeizaitcev/gophermart/internal/accrual/models"
	"github.com/sergeizaitcev/gophermart/internal/accrual/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	testOrderNumber  = "1234567890"
	testInvalidOrder = "12343245250"
	testMatchName1   = "testItem1"
	testMatchName2   = "testItem2"
)

const test_dsn = "postgresql://postgres:postgres@localhost:5432/accrual?sslmode=disable"

// testDB тест postgres
func testDB(t *testing.T) (*pgxpool.Pool, func()) {
	ctx := context.Background()

	db, err := pgxpool.New(ctx, test_dsn)
	require.NoError(t, err)
	t.Cleanup(func() { db.Close() })

	require.NoError(t, db.Ping(ctx))

	sql, err := goose.OpenDBWithDriver("pgx", test_dsn)
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, sql.Close()) })

	require.NoError(t, migrations.Up(ctx, sql))

	return db, func() { require.NoError(t, migrations.Down(ctx, sql)) }
}

func TestAccrualPostgres(t *testing.T) {
	db, teardown := testDB(t)
	defer teardown()

	accrual := NewAccrual(db)

	//Тест создания match
	_, err := accrual.CreateMatch(context.Background(), &storage.Match{MatchName: testMatchName1, Reward: 10, Type: 0})
	assert.NoError(t, err)
	_, err = accrual.CreateMatch(context.Background(), &storage.Match{MatchName: testMatchName2, Reward: 10, Type: 1})
	assert.NoError(t, err)

	//Тест на дубль match
	matchIDNil, err := accrual.CreateMatch(context.Background(), &storage.Match{MatchName: testMatchName1, Reward: 10, Type: 0})
	assert.ErrorIs(t, err, models.ErrDuplicate)
	assert.Equal(t, uuid.Nil, matchIDNil)

	//Тест получения matchID по имени
	testMatchID1, err := accrual.GetMatchByName(context.Background(), testMatchName1)
	assert.NoError(t, err)
	testMatchID2, err := accrual.GetMatchByName(context.Background(), testMatchName2)
	assert.NoError(t, err)

	//Тест отсутствия matchID
	_, err = accrual.GetMatchByName(context.Background(), "testMatchNothing")
	assert.ErrorIs(t, err, models.ErrNotFound)

	//Тест создания order
	testGoods := make([]*storage.Goods, 2)
	testGoods[0] = &storage.Goods{MatchID: testMatchID1.MatchID, Price: 12345}
	testGoods[1] = &storage.Goods{MatchID: testMatchID2.MatchID, Price: 123425}

	testOrderID, err := accrual.CreateOrderWithGoods(context.Background(), testOrderNumber, testGoods)
	assert.NoError(t, err)
	assert.NotNil(t, testOrderID)

	//Тест создания invalid order
	err = accrual.CreateInvalidOrder(context.Background(), testInvalidOrder)
	assert.NoError(t, err)

	//Тест обновления статуса и суммы вознагрождения orders
	err = accrual.UpdateOrder(context.Background(), &storage.Order{OrderID: testOrderID, Status: 3, Accrual: 500})
	assert.NoError(t, err)

	//Тест получения заказа
	want := &storage.OrderOut{OrderNumber: testOrderNumber, Status: "processed", Accrual: 500}
	order, err := accrual.GetOrderWithGoodsByNumber(context.Background(), testOrderNumber)
	assert.NoError(t, err)
	assert.Equal(t, order, want)

	//Тест получения goods
	goods, err := accrual.GetGoodsByOrderID(context.Background(), testOrderID)
	assert.NoError(t, err)
	assert.NotNil(t, goods)
}