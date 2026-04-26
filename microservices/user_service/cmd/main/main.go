package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/RSODA/wishlist/internal/api"
	"github.com/RSODA/wishlist/internal/config"
	"github.com/RSODA/wishlist/internal/migrator"
	postgres "github.com/RSODA/wishlist/internal/repository/postgres/user"
	service "github.com/RSODA/wishlist/internal/service/user"
	wishlist "github.com/RSODA/wishlist/pkg/proto/user/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

const (
	grpcAddress    = "localhost:3030"
	gatewayAddress = "localhost:5555"
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

	list, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	wishlist.RegisterUserV1Server(grpcServer, impl)
	reflection.Register(grpcServer)

	gatewayMux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(incomingHeaderMatcher),
	)

	err = wishlist.RegisterUserV1HandlerFromEndpoint(
		context.Background(),
		gatewayMux,
		grpcAddress,
		[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
	)
	if err != nil {
		log.Fatal(err)
	}

	httpServer := &http.Server{
		Addr:    gatewayAddress,
		Handler: gatewayMux,
	}

	errCh := make(chan error, 2)

	go func() {
		log.Printf("starting gRPC server on %s", grpcAddress)
		errCh <- grpcServer.Serve(list)
	}()

	go func() {
		log.Printf("starting gateway server on %s", gatewayAddress)
		errCh <- httpServer.ListenAndServe()
	}()

	err = <-errCh
	if err != nil {
		log.Fatal(err)
	}
}

func incomingHeaderMatcher(key string) (string, bool) {
	if strings.EqualFold(key, "Authorization") {
		return "authorization", true
	}

	return runtime.DefaultHeaderMatcher(key)
}
