package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/MarkoLuna/GoEmployeeCrud/pkg/utils"
	_ "github.com/lib/pq"
)

var (
	host     = utils.GetEnv("DB_HOST", "localhost")
	port     = utils.GetEnv("DB_PORT", "5432")
	user     = utils.GetEnv("DB_USER", "postgres_user")
	password = utils.GetEnv("DB_PASSWORD", "userspw")
	dbname   = utils.GetEnv("DB_NAME", "users_db")
)

var (
	db_connection *sql.DB
)

// this method connects to postgres db using the singleton pattern
func connect() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("db connection success")
	db_connection = db
}

func GetDB() *sql.DB {
	if db_connection == nil {
		connect()
	}
	return db_connection
}
