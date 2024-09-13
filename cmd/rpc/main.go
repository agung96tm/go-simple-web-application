package main

import (
	"database/sql"
	"flag"
	"github.com/agung96tm/go-simple-web-application/internal/data"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
	"os"
)

type Config struct {
	Host string
	Port string

	DB struct {
		DSN string
	}
}

type Application struct {
	Config Config
	Logger *log.Logger
	Grpc   *grpc.Server
	Models data.Models
}

func main() {
	var config Config

	flag.StringVar(&config.Host, "host", "0.0.0.0", "Host to listen on")
	flag.StringVar(&config.Port, "port", "8080", "Port to listen on")
	flag.StringVar(&config.DB.DSN, "db-dsn", "", "Database DSN")
	flag.Parse()

	logger := log.New(os.Stdout, "rpc: ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	// db
	db, err := openDB(config)
	if err != nil {
		logger.Fatal(err)
	}

	app := Application{
		Logger: logger,
		Config: config,
		Models: data.NewModels(db),
	}

	logger.Fatal(app.serve())
}

func openDB(config Config) (*sql.DB, error) {
	conn, err := sql.Open("postgres", config.DB.DSN)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}
	return conn, nil
}
