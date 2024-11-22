package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"library-api/dtos"
	utilscontext "library-api/helpers/utilsContext"
	"library-api/models"
)

type BookRepository interface {
	GetAllRepository(context.Context, string) ([]models.AuthorBook, error)
	PostBookRepository(context.Context, dtos.RequestPostBook, models.Author) (*models.AuthorBook, error)
	IsBookHasTheSameTitle(context.Context, string) bool
	GetBookByID(context.Context, int64) (*models.Book, error)
	PutQuantityBook(context.Context, int, int64) error
	IsBookOutOfStock(int64) bool
}

type BookRepositoryImplementation struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepositoryImplementation {
	return &BookRepositoryImplementation{
		db: db,
	}
}

func (br *BookRepositoryImplementation) GetAllRepository(ctx context.Context, query string) ([]models.AuthorBook, error) {
	serchKey := "%" + query + "%"
	q := `
		SELECT 
		*
		FROM books b 
		JOIN authors a ON a.id = b.author_id 
		WHERE b.deleted_at IS NULL AND title ILIKE $1
		ORDER BY b.id ASC;
	`
	rows, err := br.db.QueryContext(ctx, q, serchKey)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()
	booksAndAuthor := []models.AuthorBook{}
	for rows.Next() {
		var book models.Book
		var author models.Author
		err := rows.Scan(
			&book.ID,
			&book.AuthorId,
			&book.Title,
			&book.Description,
			&book.Quantity,
			&book.Cover,
			&book.CreatedAt,
			&book.UpdatedAt,
			&book.DeleteAt,
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
		resp := models.AuthorBook{
			Book:   book,
			Author: author,
		}
		booksAndAuthor = append(booksAndAuthor, resp)
	}
	return booksAndAuthor, nil
}

func (br *BookRepositoryImplementation) PostBookRepository(ctx context.Context, reqBody dtos.RequestPostBook, getAuthor models.Author) (*models.AuthorBook, error) {
	sql := `
		INSERT INTO Books (title, author_id, description, quantity, cover, created_at, updated_at) VALUES 
		($1, $2, $3, $4, $5, NOW(), NOW()) 
		RETURNING *;
	`
	var book models.Book
	err := br.db.QueryRowContext(ctx, sql, reqBody.Title, reqBody.AuthorId, reqBody.Description, reqBody.Quantity, reqBody.Cover).Scan(
		&book.ID,
		&book.AuthorId,
		&book.Title,
		&book.Description,
		&book.Quantity,
		&book.Cover,
		&book.CreatedAt,
		&book.UpdatedAt,
		&book.DeleteAt,
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	resp := models.AuthorBook{
		Book:   book,
		Author: getAuthor,
	}
	return &resp, nil
}

func (br *BookRepositoryImplementation) IsBookHasTheSameTitle(ctx context.Context, title string) bool {
	sql := `
		SELECT
		*
		FROM books
		WHERE title = $1 AND deleted_at IS NULL;
	`
	var book models.Book
	err := br.db.QueryRowContext(ctx, sql, title).Scan(
		&book.ID,
		&book.AuthorId,
		&book.Title,
		&book.Description,
		&book.Quantity,
		&book.Cover,
		&book.CreatedAt,
		&book.UpdatedAt,
		&book.DeleteAt,
	)
	return err != nil
}

func (br *BookRepositoryImplementation) GetBookByID(ctx context.Context, id int64) (*models.Book, error) {
	sql := `
		SELECT
		*
		FROM books
		WHERE id = $1 AND deleted_at IS NULL;
	`
	var book models.Book
	txFromCtx := utilscontext.GetTx(ctx)
	if txFromCtx != nil {
		err := txFromCtx.QueryRowContext(ctx, sql, id).Scan(
			&book.ID,
			&book.AuthorId,
			&book.Title,
			&book.Description,
			&book.Quantity,
			&book.Cover,
			&book.CreatedAt,
			&book.UpdatedAt,
			&book.DeleteAt,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	} else {
		err := br.db.QueryRowContext(ctx, sql, id).Scan(
			&book.ID,
			&book.AuthorId,
			&book.Title,
			&book.Description,
			&book.Quantity,
			&book.Cover,
			&book.CreatedAt,
			&book.UpdatedAt,
			&book.DeleteAt,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	return &book, nil
}

func (br *BookRepositoryImplementation) PutQuantityBook(ctx context.Context, newQuantity int, id int64) error {
	sql := `
		UPDATE books SET 
		quantity = $1,
		updated_at = NOW()
		WHERE id = $2;
	`
	txFromCtx := utilscontext.GetTx(ctx)
	if txFromCtx != nil {
		_, err := txFromCtx.ExecContext(ctx, sql, newQuantity, id)
		return err
	}
	_, err := br.db.ExecContext(ctx, sql, newQuantity, id)
	return err
}

func (br *BookRepositoryImplementation) IsBookOutOfStock(quantityBook int64) bool {
	return quantityBook <= 0
}
