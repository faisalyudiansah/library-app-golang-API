package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"library-api/models"
)

type AuthorRepository interface {
	GetAuthorById(context.Context, int64) (*models.Author, error)
}

type AuthorRepositoryImplementation struct {
	db *sql.DB
}

func NewAuthorRepository(db *sql.DB) *AuthorRepositoryImplementation {
	return &AuthorRepositoryImplementation{
		db: db,
	}
}

func (ar *AuthorRepositoryImplementation) GetAuthorById(ctx context.Context, id int64) (*models.Author, error) {
	sql := `
		SELECT
		*
		FROM authors 
		WHERE id = $1 AND deleted_at IS NULL;
	`
	var author models.Author
	err := ar.db.QueryRowContext(ctx, sql, id).Scan(
		&author.ID,
		&author.Name,
		&author.CreatedAt,
		&author.UpdatedAt,
		&author.DeleteAt,
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &author, nil
}
