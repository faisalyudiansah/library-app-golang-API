package servers

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"library-api/databases"

	"github.com/joho/godotenv"
)

func RunServer() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	db, err := databases.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	SetupController := SetupController(db)
	r := SetupRoute(SetupController)

	AddrConfig := os.Getenv("ADDR_CONFIG")
	srv := http.Server{
		Addr:    AddrConfig,
		Handler: r,
	}
	// StartWithoutGracefulShutdown(AddrConfig, &srv)
	StartWithGracefulShutdown(&srv)
}

func StartWithoutGracefulShutdown(AddrConfig string, srv *http.Server) {
	log.Printf("Server running on %s", AddrConfig)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("error server listen and serve: %s", err.Error())
	}
}

func StartWithGracefulShutdown(srv *http.Server) {
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	<-ctx.Done()
	log.Println("timeout of 5 seconds.")
	log.Println("Server exiting")
}
