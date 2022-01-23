package main

import (
	"context"
	"eftal/medium/mocks"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v4"
)

func TestMakeAnOrder(t *testing.T) {
	ctx := context.Background()
	req := &orderRequest{
		TableNumber: "B-1",
		Products: map[string]int{
			"Pizza":          5,
			"Chefs' Special": 3,
			"Coke":           2,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	pool := mocks.NewMockFakePoolInterface(ctrl)
	dummyError := fmt.Errorf("db.BeginTx failed with err: connection failed")
	pool.EXPECT().BeginTx(ctx, pgx.TxOptions{}).Return(nil, dummyError)

	ordersvc := &OrderService{db: pool}
	_, err := ordersvc.Make(ctx, req)
	if !errors.Is(err, dummyError) {
		t.Fatalf("erros doesn't matched!")
	}
}

func TestMakeAnOrderExecErr(t *testing.T) {
	ctx := context.Background()
	req := &orderRequest{
		TableNumber: "B-1",
		Products: map[string]int{
			"Pizza": 5,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	pool := mocks.NewMockFakePoolInterface(ctrl)
	tx := mocks.NewMockTx(ctrl)
	pool.EXPECT().BeginTx(ctx, pgx.TxOptions{}).Return(tx, nil)

	insertQuery := `INSERT INTO orders(time, table_number, product_id, product_total) VALUES($1, $2, $3)`
	execErr := fmt.Errorf("tx.Exec failed with err: failed")

	for k, v := range req.Products {
		tx.EXPECT().Exec(ctx, insertQuery, time.Now().Format(time.RFC3339), req.TableNumber, k, v).Times(1).Return(nil, execErr)
	}

	tx.EXPECT().Rollback(ctx).Times(1).Return(nil)

	ordersvc := &OrderService{db: pool}
	_, err := ordersvc.Make(ctx, req)
	if !errors.Is(err, execErr) {
		t.Fatalf("erros doesn't matched!")
	}
}
