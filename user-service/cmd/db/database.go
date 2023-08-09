package db

import (
	"database/sql"
	"fmt"

	"github.com/AthanatiusC/mandiri-miniproject/user-service/config"
	_ "github.com/lib/pq"
)

func InitDB(config config.Database, secret string) (*sql.DB, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		config.User,
		secret,
		config.Host,
		config.Port,
		config.Database,
	)
	conn, err := sql.Open("postgres", connString)
	if err != nil {
		fmt.Println("Error connecting to the database: ", err)
		return nil, err
	}
	if err := conn.Ping(); err != nil {
		fmt.Println("Error connecting to the database: ", err)
		return nil, err
	}
	return conn, nil
}
