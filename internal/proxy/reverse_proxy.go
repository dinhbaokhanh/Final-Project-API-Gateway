package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

// Timeout tối đa chờ phản hồi từ backend — tránh goroutine bị block vô hạn khi backend treo
const backendTimeout = 10 * time.Second

// NewReverseProxy tạo reverse proxy đến backend đích với timeout cứng 10 giây
func NewReverseProxy(target string) (http.Handler, error) {
	targetURL, err := url.Parse(target)
	if err != nil {
		return nil, err
	}

	p := httputil.NewSingleHostReverseProxy(targetURL)

	// Gắn http.Client có timeout để tránh backend chậm/treo làm nghẽn Gateway
	p.Transport = &http.Transport{
		ResponseHeaderTimeout: backendTimeout, // Timeout chờ phản hồi header từ backend
		MaxIdleConns:          100,            // Tối đa 100 kết nối idle trong pool
		MaxIdleConnsPerHost:   10,             // Tối đa 10 idle connection per backend host
		IdleConnTimeout:       90 * time.Second,
	}

	// Ghi đè Director để fix header Host — nếu không Go gửi "localhost:8080" thay vì host thật
	originalDirector := p.Director
	p.Director = func(req *http.Request) {
		originalDirector(req)
		req.Host = targetURL.Host
	}

	// Trả lỗi 502 thay vì để lộ thông báo lỗi kỹ thuật ra ngoài
	p.ErrorHandler = func(w http.ResponseWriter, r *http.Request, proxyErr error) {
		http.Error(w, "Dịch vụ backend hiện không khả dụng", http.StatusBadGateway)
	}

	return p, nil
}
