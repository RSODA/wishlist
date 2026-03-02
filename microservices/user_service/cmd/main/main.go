package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/RSODA/wishlist/internal/api"
	"github.com/RSODA/wishlist/internal/config"
	"github.com/RSODA/wishlist/internal/migrator"
	postgres "github.com/RSODA/wishlist/internal/repository/postgres/user"
	service "github.com/RSODA/wishlist/internal/service/user"
	wishlist "github.com/RSODA/wishlist/pkg/proto/user/v1"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.NewPostgresConfig()

	db, err := pgxpool.New(context.Background(), cfg.DSN())
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping(context.Background())
	if err != nil {
		log.Fatal("err ping database: ", err)
		return
	}

	migrationCfg := config.NewMigrationsConfig()

	migratorRunner := migrator.NewMigrator(stdlib.OpenDB(*db.Config().ConnConfig), migrationCfg.MigrationPath())

	err = migratorRunner.Up()
	if err != nil {
		log.Fatal("err migrator up: ", err)
		return
	}

	repo := postgres.NewPostgres(db)
	services := service.NewUserService(repo)
	impl := api.NewImplementation(services)

	list, err := net.Listen("tcp", "localhost:3030")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("starting server on :3030")

	server := grpc.NewServer()

	wishlist.RegisterUserV1Server(server, impl)
	reflection.Register(server)

	err = server.Serve(list)
	if err != nil {
		log.Fatal(err)
	}
}
