package utilscontext

import (
	"context"
	"database/sql"

	"library-api/models"
)

func SetTx(c context.Context, tx *sql.Tx) context.Context {
	var ctx models.Ctx = "ctx"
	return context.WithValue(c, ctx, tx)
}
