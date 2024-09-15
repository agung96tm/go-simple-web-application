package main

import (
	"database/sql"
	"flag"
	"github.com/agung96tm/go-simple-web-application/internal/data"
	_ "github.com/lib/pq"
	"html/template"
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
	Config        Config
	TemplateCache map[string]*template.Template
	Logger        *log.Logger
	Models        data.Models
}

func main() {
	var config Config

	flag.StringVar(&config.Host, "host", "0.0.0.0", "Host to listen on")
	flag.StringVar(&config.Port, "port", "8080", "Port to listen on")
	flag.StringVar(&config.DB.DSN, "db-dsn", "", "Database DSN")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Fatal(err)
	}

	db, err := openDB(config)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	app := Application{
		Config:        config,
		Logger:        logger,
		TemplateCache: templateCache,
		Models:        data.NewModels(db),
	}

	logger.Fatalln(app.serve())
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
