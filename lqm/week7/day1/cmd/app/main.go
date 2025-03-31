package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"context"
	"time"

	"github.com/tuannguyenandpadcojp/go-training/lqm/week7/day1/internal/db"
	"github.com/tuannguyenandpadcojp/go-training/lqm/week7/day1/config"
	"github.com/tuannguyenandpadcojp/go-training/lqm/week7/day1/cmd/server"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// Load .env file
	config, err := config.LoadConfig("")
	if err != nil {
		log.Fatalf("failed to read env variables")
	}

	// Construct MySQL DSN from environment variables
	dsn := config.DSN()

	// Run migrations
	m, err := migrate.New("file://migrations", "mysql://"+dsn)
	if err != nil {
		log.Fatalf("failed to initialize migrations: %v", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("failed to apply migrations: %v", err)
	}
	log.Println("Migrations applied successfully")

	// Connect to MySQL
	mysqlDB, err := db.NewMySQLDB(dsn)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer mysqlDB.Close()

	// Start gRPC server
	s := server.NewServer(config, mysqlDB)

	s.Start()

	// init signal channel
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig

	// shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s.Stop(ctx)
}