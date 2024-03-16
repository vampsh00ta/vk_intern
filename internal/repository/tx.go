package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

type Tx interface {
	Begin(ctx context.Context) context.Context
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
	getTx(ctx context.Context) *gorm.DB
}

func (db Pg) Rollback(ctx context.Context) error {
	tx := db.getTx(ctx)
	if tx == nil {
		return errors.New("no tx")
	}
	tx.Rollback()
	return nil
}
func (db Pg) Commit(ctx context.Context) error {
	tx := db.getTx(ctx)
	if tx == nil {
		return errors.New("no tx")
	}
	tx.Commit()
	return nil
}
func (db Pg) Begin(ctx context.Context) context.Context {
	tx := db.client.Begin()
	txCtx := context.WithValue(ctx, "tx", tx)
	return txCtx
}
func (db Pg) getTx(ctx context.Context) *gorm.DB {
	tx := ctx.Value("tx")
	txModel, ok := tx.(*gorm.DB)
	if !ok {
		return db.client
	} else {
		return txModel

	}
}
