package app

import (
	"fmt"
	"net/http"

	"github.com/dinhbaokhanh/ptit-gateway/internal/config"
	"github.com/dinhbaokhanh/ptit-gateway/internal/middleware"
	"github.com/dinhbaokhanh/ptit-gateway/internal/routing"
)

type App struct {
	server *http.Server
}

func New(cfg config.Config) (*App, error) {
	router, err := routing.NewRouter(cfg)
	if err != nil {
		return nil, err
	}

	handler := middleware.Chain(
		router,
		middleware.Recoverer,
		middleware.RequestLogger,
		middleware.CORS,
	)

	return &App{
		server: &http.Server{
			Addr:    fmt.Sprintf(":%s", cfg.Port),
			Handler: handler,
		},
	}, nil
}

func (a *App) Run() error {
	return a.server.ListenAndServe()
}
