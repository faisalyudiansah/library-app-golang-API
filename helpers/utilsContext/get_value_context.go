package utilscontext

import (
	"context"
	"database/sql"

	"library-api/models"
)

func GetValueUserIdFromToken(c context.Context) int64 {
	var key models.ID = "userId"
	if userId, ok := c.Value(key).(int64); ok {
		return userId
	}
	return 0
}

func GetTx(c context.Context) *sql.Tx {
	var ctx models.Ctx = "ctx"
	if tx, ok := c.Value(ctx).(*sql.Tx); ok {
		return tx
	}
	return nil
}
