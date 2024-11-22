package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/dtos"
	utilscontext "git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/helpers/utilsContext"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/models"
)

type BorrowRepository interface {
	IsUserBorrowNow(context.Context, int64) bool
	PostNewBorrow(context.Context, dtos.RequestBorrowBook, int64) (*models.Borrow, error)
	IsAlreadyReturnBook(context.Context, int64, int64, int64) bool
	IsBorrowIdValid(context.Context, int64, int64) bool
	PostReturnBook(context.Context, dtos.RequestReturnBook) (*models.Borrow, error)
}

type BorrowRepositoryImplementation struct {
	db *sql.DB
}

func NewBorrowRepository(db *sql.DB) *BorrowRepositoryImplementation {
	return &BorrowRepositoryImplementation{
		db: db,
	}
}

func (brr *BorrowRepositoryImplementation) IsUserBorrowNow(ctx context.Context, userId int64) bool {
	sql := `
		SELECT
		*
		FROM borrows
		WHERE user_id = $1 AND return_date IS NULL AND deleted_at IS NULL;
	`
	var borrow models.Borrow
	txFromCtx := utilscontext.GetTx(ctx)
	if txFromCtx != nil {
		err := txFromCtx.QueryRowContext(ctx, sql, userId).Scan(
			&borrow.ID,
			&borrow.UserID,
			&borrow.BookID,
			&borrow.BorrowDate,
			&borrow.ReturnDate,
			&borrow.CreatedAt,
			&borrow.UpdatedAt,
			&borrow.DeleteAt,
		)
		return err == nil
	} else {
		err := brr.db.QueryRowContext(ctx, sql, userId).Scan(
			&borrow.ID,
			&borrow.UserID,
			&borrow.BookID,
			&borrow.BorrowDate,
			&borrow.ReturnDate,
			&borrow.CreatedAt,
			&borrow.UpdatedAt,
			&borrow.DeleteAt,
		)
		return err == nil
	}
}

func (brr *BorrowRepositoryImplementation) PostNewBorrow(ctx context.Context, reqBody dtos.RequestBorrowBook, userId int64) (*models.Borrow, error) {
	sql := `
		INSERT INTO borrows (user_id, book_id, borrow_date, created_at, updated_at) VALUES 
		($1, $2, $3, NOW(), NOW())
		RETURNING *;
	`
	var borrow models.Borrow
	txFromCtx := utilscontext.GetTx(ctx)
	if txFromCtx != nil {
		err := txFromCtx.QueryRowContext(ctx, sql, userId, reqBody.BookID, reqBody.BorrowDate).Scan(
			&borrow.ID,
			&borrow.UserID,
			&borrow.BookID,
			&borrow.BorrowDate,
			&borrow.ReturnDate,
			&borrow.CreatedAt,
			&borrow.UpdatedAt,
			&borrow.DeleteAt,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	} else {
		err := brr.db.QueryRowContext(ctx, sql, userId, reqBody.BookID, reqBody.BorrowDate).Scan(
			&borrow.ID,
			&borrow.UserID,
			&borrow.BookID,
			&borrow.BorrowDate,
			&borrow.ReturnDate,
			&borrow.CreatedAt,
			&borrow.UpdatedAt,
			&borrow.DeleteAt,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	return &borrow, nil
}

func (brr *BorrowRepositoryImplementation) IsAlreadyReturnBook(ctx context.Context, id int64, userId int64, bookId int64) bool {
	sql := `
		SELECT
		*
		FROM borrows
		WHERE id = $1 AND user_id = $2 AND book_id = $3 AND return_date IS NOT NULL AND deleted_at IS NULL;
	`
	var borrow models.Borrow
	txFromCtx := utilscontext.GetTx(ctx)
	if txFromCtx != nil {
		txFromCtx.QueryRowContext(ctx, sql, id, userId, bookId).Scan(
			&borrow.ID,
			&borrow.UserID,
			&borrow.BookID,
			&borrow.BorrowDate,
			&borrow.ReturnDate,
			&borrow.CreatedAt,
			&borrow.UpdatedAt,
			&borrow.DeleteAt,
		)
	} else {
		brr.db.QueryRowContext(ctx, sql, id, userId, bookId).Scan(
			&borrow.ID,
			&borrow.UserID,
			&borrow.BookID,
			&borrow.BorrowDate,
			&borrow.ReturnDate,
			&borrow.CreatedAt,
			&borrow.UpdatedAt,
			&borrow.DeleteAt,
		)
	}
	return borrow.ID != 0
}

func (brr *BorrowRepositoryImplementation) IsBorrowIdValid(ctx context.Context, borrowId int64, bookId int64) bool {
	sql := `
		SELECT
		*
		FROM borrows
		WHERE id = $1 AND book_id = $2 AND deleted_at IS NULL;
	`
	var borrow models.Borrow
	txFromCtx := utilscontext.GetTx(ctx)
	if txFromCtx != nil {
		txFromCtx.QueryRowContext(ctx, sql, borrowId, bookId).Scan(
			&borrow.ID,
			&borrow.UserID,
			&borrow.BookID,
			&borrow.BorrowDate,
			&borrow.ReturnDate,
			&borrow.CreatedAt,
			&borrow.UpdatedAt,
			&borrow.DeleteAt,
		)
	} else {
		brr.db.QueryRowContext(ctx, sql, borrowId).Scan(
			&borrow.ID,
			&borrow.UserID,
			&borrow.BookID,
			&borrow.BorrowDate,
			&borrow.ReturnDate,
			&borrow.CreatedAt,
			&borrow.UpdatedAt,
			&borrow.DeleteAt,
		)

	}
	return borrow.ID != 0
}

func (brr *BorrowRepositoryImplementation) PostReturnBook(ctx context.Context, reqBody dtos.RequestReturnBook) (*models.Borrow, error) {
	sql := `
		UPDATE borrows SET
		return_date = $1,
		updated_at = NOW()
		WHERE id = $2 AND deleted_at IS NULL
		RETURNING *;
	`
	var borrow models.Borrow
	txFromCtx := utilscontext.GetTx(ctx)
	if txFromCtx != nil {
		err := txFromCtx.QueryRowContext(ctx, sql, reqBody.ReturnDate, reqBody.BorrowID).Scan(
			&borrow.ID,
			&borrow.UserID,
			&borrow.BookID,
			&borrow.BorrowDate,
			&borrow.ReturnDate,
			&borrow.CreatedAt,
			&borrow.UpdatedAt,
			&borrow.DeleteAt,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	} else {
		err := brr.db.QueryRowContext(ctx, sql, reqBody.ReturnDate, reqBody.BorrowID).Scan(
			&borrow.ID,
			&borrow.UserID,
			&borrow.BookID,
			&borrow.BorrowDate,
			&borrow.ReturnDate,
			&borrow.CreatedAt,
			&borrow.UpdatedAt,
			&borrow.DeleteAt,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

	}
	return &borrow, nil
}
