package repositories

import (
	"context"
	"database/sql"
	"fmt"

	utilscontext "git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/helpers/utilsContext"
)

type TransactionRepository interface {
	Atomic(c context.Context, fn func(context.Context) (any, error)) (any, error)
}

type TransactionRepositoryImpelementation struct {
	db *sql.DB
}

func NewTransactionRepositoryImpelementation(db *sql.DB) *TransactionRepositoryImpelementation {
	return &TransactionRepositoryImpelementation{
		db: db,
	}
}

func (dc *TransactionRepositoryImpelementation) Atomic(c context.Context, fn func(context.Context) (any, error)) (any, error) {
	tx, err := dc.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				err = fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
			}
		}
		err = tx.Commit()
	}()

	result, err := fn(utilscontext.SetTx(c, tx))
	if err != nil {
		return nil, err
	}
	return result, nil
}
