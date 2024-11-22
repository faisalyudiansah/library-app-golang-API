package servers

import (
	"library-api/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoute(h *HandlerOps) *gin.Engine {
	g := gin.New()
	g.ContextWithFallback = true
	g.Use(middlewares.IncrementRequestCount, middlewares.LoggerMiddleware(), gin.Recovery(), middlewares.ErrorHandler)

	g.NoRoute(InvalidRoute)

	SetupAuthenRoutes(g, h)
	SetupBookRoutes(g, h)
	SetupBorrowRoutes(g, h)

	return g
}

func SetupAuthenRoutes(g *gin.Engine, h *HandlerOps) {
	g.POST("/register", h.UserController.PostRegisterUserController)
	g.POST("/login", h.UserController.PostLoginUserController)
}

func SetupBookRoutes(g *gin.Engine, h *HandlerOps) {
	g.GET("/books", h.BookController.GetAllBookController)
	g.POST("/books", h.BookController.PostBookController)
}

func SetupBorrowRoutes(g *gin.Engine, h *HandlerOps) {
	g.POST("/borrow-books", middlewares.AuthorizationBorrow, h.BorrowController.PostNewBorrowBookController)
	g.POST("/return-books", middlewares.AuthorizationBorrow, h.BorrowController.PostReturnBookController)
}
