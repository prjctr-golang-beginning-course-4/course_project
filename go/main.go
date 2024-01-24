package main

import (
	"context"
	"course/src"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/pressly/goose"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ctx = context.Background()

func main() {
	//err := godotenv.Load("../.env")
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Failed to load .env file: %v\n", err)
	}

	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlDB := os.Getenv("MYSQL_DATABASE")

	//dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", mysqlUser, mysqlPassword, mysqlDB)
	dsn := fmt.Sprintf("%s:%s@tcp(db:3306)/%s", mysqlUser, mysqlPassword, mysqlDB)
	redisPassword := os.Getenv("REDIS_PASSWORD")

	db, err := connectDB(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to db: %v\n", err)
	}
	defer db.Close()

	err = goose.SetDialect("mysql")
	if err != nil {
		log.Fatalf("Failed to set DB dialect: %v\n", err)
	}

	err = goose.Up(db, "src/migrations")
	if err != nil {
		log.Fatalf("Failed to run migrations: %v\n", err)
	}

	gorm, err := gorm.Open(mysql.Open(dsn + "?charset=utf8&parseTime=True&loc=Local"))
	if err != nil {
		log.Fatalf("Failed to initialize gorm: %v\n", err)
	}

	rdb, err := connectCache(redisPassword)
	if err != nil {
		log.Fatalf("Failed to connect to redis: %v\n", err)
	}

	src.Run(gorm, rdb)
}

func connectDB(dbSource string) (*sql.DB, error) {
	var db *sql.DB
	var err error

	for i := 0; i < 5; i++ {
		db, err = sql.Open("mysql", dbSource)

		if err != nil {
			fmt.Println("Error connecting to the database. Retrying...")

			time.Sleep(time.Second * 2)
			continue
		}

		if err = db.Ping(); err == nil {
			return db, nil
		}
	}

	return nil, err
}

func connectCache(cacheSource string) (*redis.Client, error) {
	var rdb *redis.Client
	var err error
	var pong string

	for i := 0; i < 5; i++ {
		rdb = redis.NewClient(&redis.Options{
			//Addr:     "localhost:6379",
			Addr:     "cache:6379",
			Password: cacheSource,
			DB:       0,
		})

		pong, err = rdb.Ping(ctx).Result()

		if err != nil || pong != "PONG" {
			fmt.Println("Error connecting to Redis. Retrying...")

			time.Sleep(time.Second * 2)
			continue
		}

		if err == nil {
			return rdb, nil
		}
	}

	return nil, err
}
