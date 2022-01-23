package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type OrderService struct {
	db FakePoolInterface
}

type OrderInterface interface {
	Make(ctx context.Context, req *orderRequest) (*orderResponse, error)
}

type orderResponse struct{}

func NewOrder(db *pgxpool.Pool) OrderInterface {
	return &OrderService{
		db: db,
	}
}

type orderRequest struct {
	TableNumber string         `json:"table_number"`
	Products    map[string]int `json:"products"`
}

// Make checks orderRequest and write to the orders table
func (o *OrderService) Make(ctx context.Context, req *orderRequest) (*orderResponse, error) {
	if len(req.TableNumber) == 0 {
		return nil, errors.New("tableNumber varible is not set")
	}

	if len(req.Products) == 0 {
		return nil, errors.New("orders variable is not set")
	}

	tx, err := o.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("db.BeginTx failed with error: %w", err)
	}

	insertQuery := `INSERT INTO orders(time, table_number, product_id, product_total) VALUES($1, $2, $3)`

	for k, v := range req.Products {
		_, err := tx.Exec(ctx, insertQuery, time.Now().Format(time.RFC3339), req.TableNumber, k, v)
		if err != nil {
			tx.Rollback(ctx)
			return nil, fmt.Errorf("tx.Exec failed with error: %w", err)
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("tx.Commit failed with err: %w", err)
	}

	return &orderResponse{}, nil
}
