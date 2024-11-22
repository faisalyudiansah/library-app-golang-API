package dtos

import "library-api/models"

func ToResponseBookType(books []models.Book) []ResponseBook {
	var bookResponses []ResponseBook
	for _, book := range books {
		bookResponses = append(bookResponses, ToDtoResponseBook(book))
	}
	return bookResponses
}

func ToResponseBookAuthor(books []models.AuthorBook) []ResponseBook {
	var bookResponses []ResponseBook
	for _, book := range books {
		bookResponses = append(bookResponses, ToDtoResponseBookAuthor(book))
	}
	return bookResponses
}

func ToDtoResponseBookAuthor(bookData models.AuthorBook) ResponseBook {
	return ResponseBook{
		ID:          bookData.Book.ID,
		AuthorId:    bookData.Book.AuthorId,
		Title:       bookData.Book.Title,
		Description: bookData.Book.Description,
		Quantity:    bookData.Book.Quantity,
		Cover:       bookData.Book.Cover,
		CreatedAt:   bookData.Book.CreatedAt,
		UpdatedAt:   bookData.Book.UpdatedAt,
		DeleteAt:    bookData.Book.DeleteAt,
		Author:      ResponseAuthor(bookData.Author),
	}
}

func ToDtoResponseBook(book models.Book) ResponseBook {
	return ResponseBook{
		ID:          book.ID,
		Title:       book.Title,
		Description: book.Description,
		Quantity:    book.Quantity,
		Cover:       book.Cover,
		CreatedAt:   book.CreatedAt,
		UpdatedAt:   book.UpdatedAt,
		DeleteAt:    book.DeleteAt,
	}
}

func ToDtoResponseBorrow(borrow *models.Borrow) *ResponseBorrow {
	return &ResponseBorrow{
		ID:         borrow.ID,
		UserID:     borrow.UserID,
		BookID:     borrow.BookID,
		BorrowDate: borrow.BorrowDate,
		ReturnDate: borrow.ReturnDate,
		CreatedAt:  borrow.CreatedAt,
		UpdatedAt:  borrow.UpdatedAt,
		DeleteAt:   borrow.DeleteAt,
	}
}

func ToDtoResponseUserInfo(user *models.User) ResponseDataUser {
	return ResponseDataUser{
		Name:      &user.Name,
		Email:     &user.Email,
		CreatedAt: &user.CreatedAt,
	}
}

func ToDtoResponseUserAccessToken(ac string) ResponseDataUser {
	return ResponseDataUser{
		AccessToken: &ac,
	}
}
