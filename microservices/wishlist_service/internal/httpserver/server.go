package httpserver

import (
	"context"
	"net/http"

	wish_v1 "github.com/RSODA/wishlist/microservices/wishlist_service/pkg/proto/wish/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func NewServer(addr string, wishServer wish_v1.WishV1Server) (*http.Server, error) {
	gatewayMux := runtime.NewServeMux()
	if err := wish_v1.RegisterWishV1HandlerServer(context.Background(), gatewayMux, wishServer); err != nil {
		return nil, err
	}

	rootMux := http.NewServeMux()
	rootMux.Handle("/api/", gatewayMux)
	rootMux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	rootMux.Handle("/api/v1/static/", http.StripPrefix(
		"/api/v1/static/",
		http.FileServer(http.Dir("assets/img")),
	))

	return &http.Server{
		Addr:    addr,
		Handler: withCORS(rootMux),
	}, nil
}

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			origin = "*"
		}

		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "false")
		w.Header().Set("Vary", "Origin")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
