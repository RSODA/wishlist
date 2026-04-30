package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	wish_api "github.com/RSODA/wishlist/microservices/wishlist_service/internal/api/wish"
	"github.com/RSODA/wishlist/microservices/wishlist_service/internal/client/user"
	"github.com/RSODA/wishlist/microservices/wishlist_service/internal/config"
	"github.com/RSODA/wishlist/microservices/wishlist_service/internal/httpserver"
	"github.com/RSODA/wishlist/microservices/wishlist_service/internal/migrator"
	repo_wish "github.com/RSODA/wishlist/microservices/wishlist_service/internal/repository/postgres/wish"
	service_wish "github.com/RSODA/wishlist/microservices/wishlist_service/internal/service/wish"
	wish_v1 "github.com/RSODA/wishlist/microservices/wishlist_service/pkg/proto/wish/v1"
	userv1 "github.com/RSODA/wishlist/pkg/proto/user/v1"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, os.Interrupt, syscall.SIGTERM)

	config.Load()

	httpAddr, err := config.NewHTTPConfig()
	if err != nil {
		log.Fatal("err init http config: ", err)
		return
	}

	dsn, err := config.NewPGConfig()
	if err != nil {
		log.Fatal("err init database config: ", err)
	}

	grpcCfg := config.NewGRPCConfig()
	migrationpath, err := config.NewMigrationsConfig()
	if err != nil {
		log.Fatal("err init migrations config: ", err)
	}

	db, err := pgxpool.New(context.Background(), dsn.DSN())
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping(context.Background())
	if err != nil {
		log.Fatal("err ping database: ", err)
		return
	}

	mig := migrator.NewMigrator(stdlib.OpenDB(*db.Config().ConnConfig), migrationpath.MigrationPath())

	err = mig.Up()
	if err != nil {
		log.Fatal(err)
	}

	userConn, err := grpc.NewClient(
		grpcCfg.UserServiceAddress(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal("err connect user_service grpc: ", err)
	}
	defer func() {
		if closeErr := userConn.Close(); closeErr != nil {
			log.Printf("err closing user_service grpc connection: %v", closeErr)
		}
	}()

	userClient := userv1.NewUserV1Client(userConn)
	userService := user.NewUserService(userClient)

	log.Printf("user_service grpc client connected to %s", grpcCfg.UserServiceAddress())

	repository, err := repo_wish.New(db)
	if err != nil {
		log.Fatal("err init repository: ", err)
	}

	wishService := service_wish.NewWishService(repository, userService, httpAddr.Address())
	wishImpl := wish_api.NewImplementation(wishService)

	server := grpc.NewServer()
	wish_v1.RegisterWishV1Server(server, wishImpl)
	reflection.Register(server)

	grpcAddr := grpcCfg.ServiceAddress()
	listener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatal("err listen grpc server: ", err)
	}

	httpServer, err := httpserver.NewServer(httpAddr.Address(), wishImpl)
	if err != nil {
		log.Fatal("err init http server: ", err)
	}

	go func() {
		log.Printf("wishlist_service grpc server started on %s", grpcAddr)
		if serveErr := server.Serve(listener); serveErr != nil {
			log.Fatalf("err serve grpc server: %v", serveErr)
		}
	}()

	go func() {
		log.Printf("wishlist_service http server started on %s", grpcCfg.HTTPAddress())
		if serveErr := httpServer.ListenAndServe(); serveErr != nil && serveErr != http.ErrServerClosed {
			log.Fatalf("err serve http server: %v", serveErr)
		}
	}()

	<-shutdownCh
	log.Println("shutdown signal received")
	server.GracefulStop()
	if err = httpServer.Shutdown(context.Background()); err != nil {
		log.Printf("err shutdown http server: %v", err)
	}
	db.Close()
	log.Println("wishlist_service stopped")

}
