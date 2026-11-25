package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	// httpSwagger "github.com/swaggo/http-swagger/v2"
	"github.com/uforco/sqlitebygo/cmd/api/registar-api"
	// _ "github.com/uforco/sqlitebygo/docs"
	"github.com/uforco/sqlitebygo/internal/config"
	"github.com/uforco/sqlitebygo/internal/db/sqlite"
)

func main() {
	// load conig
	cfg := config.MustLoad(".env")

	// database connection
	db, dbError := sqlite.New(cfg)
	if dbError != nil {
		log.Fatal("failed to init db", dbError.Error())
	}
	
	// setup router
	router := http.NewServeMux()

	// setup handlers
	router.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	// register static files
	registar.Api(router, cfg, db)

	// setup swag
	// router.Handle("/swagger/", httpSwagger.Handler(
	// 	httpSwagger.URL("/swagger/doc.json"),
	// ))

	// setup server
	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}
	// concarece cal for server start and stop
	// Try to Start server
	slog.Info("server started", "address", cfg.Address)
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("failed to start server")
		}
	}()
	<-done
	// if Err and user close server; then shutdown server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}
	slog.Info("server stopped")

}
