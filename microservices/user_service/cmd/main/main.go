package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/RSODA/wishlist/internal/api"
	"github.com/RSODA/wishlist/internal/config"
	"github.com/RSODA/wishlist/internal/interceptor"
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

func main() {
	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, os.Interrupt, syscall.SIGTERM)

	err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := config.NewPostgresConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := pgxpool.New(context.Background(), cfg.DSN())
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping(context.Background())
	if err != nil {
		log.Fatal("err ping database: ", err)
		return
	}

	migrationCfg, err := config.NewMigrationsConfig()
	if err != nil {
		log.Fatal(err)
	}

	migratorRunner := migrator.NewMigrator(stdlib.OpenDB(*db.Config().ConnConfig), migrationCfg.MigrationPath())

	err = migratorRunner.Up()
	if err != nil {
		log.Fatal("err migrator up: ", err)
		return
	}

	repo := postgres.NewPostgres(db)
	services := service.NewUserService(repo)
	impl := api.NewImplementation(services)

	grpcCfg, err := config.NewGRPCConfig()
	if err != nil {
		log.Fatal(err)
	}

	list, err := net.Listen("tcp", grpcCfg.ServiceAddress())
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.AuthInterceptor),
	)

	wishlist.RegisterUserV1Server(grpcServer, impl)
	reflection.Register(grpcServer)

	gatewayMux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(incomingHeaderMatcher),
	)

	err = wishlist.RegisterUserV1HandlerFromEndpoint(
		context.Background(),
		gatewayMux,
		grpcCfg.ServiceAddress(),
		[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
	)
	if err != nil {
		log.Fatal(err)
	}

	httpCfg, err := config.NewHTTPConfig()
	if err != nil {
		log.Fatal(err)
	}

	httpServer := &http.Server{
		Addr:    httpCfg.Address(),
		Handler: gatewayMux,
	}

	errCh := make(chan error, 2)

	go func() {
		log.Printf("starting gRPC server on %s", grpcCfg.ServiceAddress())
		errCh <- grpcServer.Serve(list)
	}()

	go func() {
		log.Printf("starting gateway server on %s", httpCfg.Address())
		errCh <- httpServer.ListenAndServe()
	}()

	select {
	case err = <-errCh:
		log.Printf("server error: %v", err)
	case <-shutdownCh:
		log.Println("shutdown signal received")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	log.Println("shutting down server gracefully")

	log.Println("shutting down gRPC server")
	grpcServer.GracefulStop()

	log.Println("shutting down HTTP server")
	httpServer.Shutdown(ctx)

	log.Println("close connect to postgres")
	db.Close()
}

func incomingHeaderMatcher(key string) (string, bool) {
	if strings.EqualFold(key, "Authorization") {
		return "authorization", true
	}

	return runtime.DefaultHeaderMatcher(key)
}
