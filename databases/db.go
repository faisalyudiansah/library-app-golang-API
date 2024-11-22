package databases

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectDB() (*sql.DB, error) {
	var (
		host     = os.Getenv("HOST_DB")
		port     = os.Getenv("PORT_DB")
		user     = os.Getenv("USER_DB")
		password = os.Getenv("PASSWORD_DB")
		dbname   = os.Getenv("DBNAME_DB")
	)
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
		host, port, user, password, dbname)
	db, err := sql.Open("pgx", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
