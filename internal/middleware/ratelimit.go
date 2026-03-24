package middleware

import (
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

// visitor lưu rate limiter và thời điểm lần cuối được thấy của một IP
type visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var (
	visitors = make(map[string]*visitor)
	mu       sync.Mutex
)

// InitRateLimiter khởi động goroutine dọn dẹp IP không hoạt động mỗi 5 phút.
// Tránh memory leak khi hàng nghìn IP unique tích lũy không được giải phóng.
func InitRateLimiter() {
	go cleanupVisitors()
}

// cleanupVisitors xóa các IP không còn gửi request trong vòng 10 phút
func cleanupVisitors() {
	for {
		time.Sleep(5 * time.Minute)

		mu.Lock()
		for ip, v := range visitors {
			if time.Since(v.lastSeen) > 10*time.Minute {
				delete(visitors, ip)
			}
		}
		mu.Unlock()
	}
}

// getVisitor trả về rate limiter của IP, tạo mới nếu chưa tồn tại
func getVisitor(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	v, exists := visitors[ip]
	if !exists {
		// Cho phép tối đa 20 request/giây, burst tối đa 20
		v = &visitor{
			limiter: rate.NewLimiter(rate.Limit(20), 20),
		}
		visitors[ip] = v
	}

	// Cập nhật thời gian hoạt động để cleanup goroutine không xóa nhầm
	v.lastSeen = time.Now()
	return v.limiter
}

// RateLimitMiddleware chặn request khi IP vượt quá 20 req/giây
func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			ip = r.RemoteAddr
		}

		if !getVisitor(ip).Allow() {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
