package main

import (
	_ "github.com/lib/pq"
	"log/slog"
	"mizar/internal/config"
	"mizar/pkg/logger"
	"mizar/pkg/server"
	"mizar/pkg/storage/postrges"
	"os"
)

func main() {
	cfg := config.Load()

	log := logger.SetupLogger(cfg.Env)
	log.Info("starting mizar", slog.String("env", cfg.Env))

	storage, err := postrges.New(cfg)
	if err != nil {
		log.Error("failed to init storage:", err)
		os.Exit(1)
	}

	srv := server.NewServer(cfg, storage)
	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}

	log.Error("server stopped")
}
