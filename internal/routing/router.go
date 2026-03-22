package routing

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dinhbaokhanh/Final-Project-API-Gateway/internal/config"
	"github.com/dinhbaokhanh/Final-Project-API-Gateway/internal/proxy"
)

// NewRouter khởi tạo một HTTP Handler để xử lý định tuyến dựa trên file cấu hình
func NewRouter(cfg *config.GatewayConfig) (http.Handler, error) {
	mux := http.NewServeMux()

	// Khai báo route kiểm tra Gateway
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	// Duyệt qua tất cả các endpoint trong cấu hình để tạo route động
	for _, endpoint := range cfg.Endpoints {
		if len(endpoint.Backend) == 0 || len(endpoint.Backend[0].Host) == 0 {
			continue 
		}

		targetURL := endpoint.Backend[0].Host[0]
		
		// Khởi tạo một bộ trung chuyển (Reverse Proxy) trỏ tới backend đích
		reverseProxy, err := proxy.NewReverseProxy(targetURL)
		if err != nil {
			return nil, fmt.Errorf("URL backend không hợp lệ cho endpoint %s: %w", endpoint.Endpoint, err)
		}

		// Tận dụng tính năng routing tiên tiến của Go 1.22+: Khai báo rõ HTTP Method (Ví dụ "POST /api/v1/user/login")
		pattern := endpoint.Endpoint
		if endpoint.Method != "" && endpoint.Method != "ANY" {
			pattern = fmt.Sprintf("%s %s", strings.ToUpper(endpoint.Method), endpoint.Endpoint)
		}

		// Đăng ký route vào bộ định tuyến
		fmt.Printf("[Router] Đã đăng ký %-30s -> chuyển hướng sang %s\n", pattern, targetURL)
		
		// Go's ReverseProxy sẽ tự động giữ nguyên đường dẫn (r.URL.Path) và gắn nó vào đằng sau TargetURL.
		// VD: Khách gọi POST /api/v1/chat/123 -> Sẽ forward đúng POST {targetURL}/api/v1/chat/123
		mux.Handle(pattern, reverseProxy)
	}

	return mux, nil
}
