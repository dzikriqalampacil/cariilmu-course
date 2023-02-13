package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

func TestEmpty(t *testing.T) {

}

func GetConnection() *sql.DB {

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	host := os.Getenv("host")
	port := os.Getenv("dbport")
	user := os.Getenv("user")
	password := os.Getenv("password")
	dbname := os.Getenv("dbname")

	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

func TestEnv(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	val := os.Getenv("host")
	fmt.Println(val)

	val = os.Getenv("dbport")
	fmt.Println(val)

	val = os.Getenv("user")
	fmt.Println(val)

	val = os.Getenv("password")
	fmt.Println(val)

	val = os.Getenv("dbname")
	fmt.Println(val)
}

func TestDB(t *testing.T) {
	db := GetConnection()
	defer db.Close()
}

func TestRouter(t *testing.T) {
	router := httprouter.New()

	server := http.Server{
		Handler: router,
		Addr:    "localhost:8000",
	}

	server.ListenAndServe()

}
