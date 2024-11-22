package constants

const (
	RequestBodyInvalid       = "request body invalid or missing"
	FailedAddProduct         = "failed to add a product"
	FailedTitleAlreadyExists = "title already exists"
	InvalidAuthorId          = "author id does not exists"
	InvalidUserId            = "user id does not exists"
	InvalidBookId            = "book id does not exists"
	InvalidBorrowIdBookId    = "borrow id / book id is not valid"
	OutOfStockBook           = "book is out of stock"
	UserAlreadyBorrowBook    = "user already borrow a book"
	UserAlreadyReturnBook    = "user already return a book"
	FailedBorrowBook         = "there was an error in the borrow process, try again"
	FailedReturnBook         = "there was an error in the return process, try again"
)

const (
	UserInvalidEmailPassword = "invalid email / password"
	UserFailedRegister       = "there was an error in the register process, try again"
	UserFailedLogin          = "there was an error in the login process, try again"
	UserEmailAlreadyExists   = "email already exists"
)

const (
	ISE                = "internal server error"
	InvalidAccessToken = "invalid access token"
	Unauthorization    = "unauthorization"
	UrlNotFound        = "url not found"
)
