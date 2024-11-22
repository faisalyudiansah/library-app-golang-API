package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"library-api/dtos"
	utilscontext "library-api/helpers/utilsContext"
	"library-api/models"
)

type UserRepository interface {
	GetUserById(context.Context, int64) (*models.User, error)
	GetUserByEmail(context.Context, string) (*models.User, error)
	PostUser(context.Context, dtos.RequestRegisterUser, string) (*models.User, error)
	IsEmailAlreadyRegistered(context.Context, string) bool
}

type UserRepositoryImplementation struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepositoryImplementation {
	return &UserRepositoryImplementation{
		db: db,
	}
}

func (brr *UserRepositoryImplementation) GetUserById(ctx context.Context, userId int64) (*models.User, error) {
	sql := `
		SELECT
		*
		FROM users
		WHERE id = $1 AND deleted_at IS NULL;
	`
	var user models.User

	txFromCtx := utilscontext.GetTx(ctx)
	if txFromCtx != nil {
		err := txFromCtx.QueryRowContext(ctx, sql, userId).Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeleteAt,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	} else {
		err := brr.db.QueryRowContext(ctx, sql, userId).Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeleteAt,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	return &user, nil
}

func (brr *UserRepositoryImplementation) GetUserByEmail(ctx context.Context, emailInput string) (*models.User, error) {
	sql := `
		SELECT
		*
		FROM users
		WHERE email = $1 AND deleted_at IS NULL;
	`
	var user models.User
	err := brr.db.QueryRowContext(ctx, sql, emailInput).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeleteAt,
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &user, nil
}

func (brr *UserRepositoryImplementation) IsEmailAlreadyRegistered(ctx context.Context, emailInput string) bool {
	sql := `
		SELECT
		id,
		name,
		email,
		password,
		created_at
		FROM users
		WHERE email = $1 AND deleted_at IS NULL;
	`
	var user models.User
	brr.db.QueryRowContext(ctx, sql, emailInput).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	return user.ID != 0
}

func (brr *UserRepositoryImplementation) PostUser(ctx context.Context, reqBody dtos.RequestRegisterUser, hashPassword string) (*models.User, error) {
	sql := `
		INSERT INTO users (name, email, password, created_at, updated_at) VALUES 
		($1, $2, $3, NOW(), NOW())
		RETURNING *;
	`
	var user models.User
	err := brr.db.QueryRowContext(ctx, sql, reqBody.Name, reqBody.Email, hashPassword).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeleteAt,
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &user, nil
}
