package routing

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dinhbaokhanh/Final-Project-API-Gateway/internal/config"
	"github.com/dinhbaokhanh/Final-Project-API-Gateway/internal/proxy"
)

func NewRouter(cfg config.Config) (http.Handler, error) {
	usersProxy, err := proxy.NewReverseProxy(cfg.UsersBackend)
	if err != nil {
		return nil, fmt.Errorf("invalid users backend URL: %w", err)
	}

	ordersProxy, err := proxy.NewReverseProxy(cfg.OrdersBackend)
	if err != nil {
		return nil, fmt.Errorf("invalid orders backend URL: %w", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	mux.Handle("/api/users/", stripPrefixThenProxy("/api/users", usersProxy))
	mux.Handle("/api/orders/", stripPrefixThenProxy("/api/orders", ordersProxy))

	return mux, nil
}

func stripPrefixThenProxy(prefix string, target http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimPrefix(r.URL.Path, prefix)
		if r.URL.Path == "" {
			r.URL.Path = "/"
		}
		target.ServeHTTP(w, r)
	})
}
