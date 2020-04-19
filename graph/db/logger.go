package db

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v9"
)

type Logs struct{}

func (l Logs) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (l Logs) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	fmt.Println(q.FormattedQuery())
	return nil
}
