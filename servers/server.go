package servers

import (
	"database/sql"

	"git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/controllers"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/helpers"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/helpers/logger"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/repositories"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/services"
)

type HandlerOps struct {
	BookController   *controllers.BookController
	BorrowController *controllers.BorrowController
	UserController   *controllers.UserController
}

func SetupController(db *sql.DB) *HandlerOps {
	logrusLogger := logger.NewLogger()
	logger.SetLogger(logrusLogger)

	bookRepository := repositories.NewBookRepository(db)
	authorRepository := repositories.NewAuthorRepository(db)
	borrowRepository := repositories.NewBorrowRepository(db)
	userRepository := repositories.NewUserRepository(db)
	transactionRepository := repositories.NewTransactionRepositoryImpelementation(db)
	bcryptStruct := helpers.NewBcryptStruct()
	jwt := helpers.NewJWTProviderHS256()

	bookService := services.NewBookServiceImplementation(bookRepository, authorRepository, borrowRepository, userRepository, transactionRepository)
	borrowService := services.NewBorrowServiceImplementation(bookRepository, borrowRepository, userRepository, transactionRepository)
	userService := services.NewUserServiceImplementation(userRepository, bcryptStruct, jwt)

	bookController := controllers.NewBookController(bookService)
	borrowController := controllers.NewBorrowController(borrowService)
	userContoller := controllers.NewUserController(userService)
	return &HandlerOps{
		BookController:   bookController,
		BorrowController: borrowController,
		UserController:   userContoller,
	}
}
