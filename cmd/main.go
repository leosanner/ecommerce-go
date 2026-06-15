package main

import (
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	cfg := config{
		addr: ":8080",
		db: dbConfig{
			dsn: "host=localhost user=postgres password=postgres dbname=ecom sslmode=disabled",
		},
	}

	conn, err := pgx.Connect(ctx, cfg.db.dsn)

	api := application{
		config: cfg,
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := api.run(api.mount()); err != nil {
		slog.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}
