package middleware

import (
	"log"
	"net/http"
	"time"
)

// Middleware định nghĩa một lớp bọc HTTP Handler, giúp thực hiện logic chặn lọc trước khi vào Server
type Middleware func(http.Handler) http.Handler

// Chain là hàm gom tất cả middleware lại và bọc vào nhau như cấu trúc củ hành tây
// Request đi từ ngoài vào trong, lớp cuối cùng sẽ là Router thực sự
func Chain(handler http.Handler, middlewares ...Middleware) http.Handler {
	wrapped := handler
	for i := len(middlewares) - 1; i >= 0; i-- {
		wrapped = middlewares[i](wrapped)
	}
	return wrapped
}

// RequestLogger: Ghi nhật ký mọi request đi qua API Gateway để xem thời gian xử lý nhanh hay chậm
func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r) // Cho phép request trôi qua để thực thi
		
		// In ra Method, Đường dẫn, và Thời gian xử lý (VD: GET /api/users 12ms)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}

// Recoverer: Đóng vai trò làm "Áo giáp bảo vệ". Nếu các hàm bên trong xảy ra lỗi nghiêm trọng (Panic), 
// cái bọc này sẽ hứng lại vớt vát để Gateway không bị sập (Crash 100%)
func Recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("Phát hiện lỗi nghiêm trọng (Panic recovered): %v", rec)
				http.Error(w, "Lỗi máy chủ nội bộ (Internal Server Error)", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

// CORS: Chia sẻ tài nguyên nguồn gốc chéo. Cảnh báo cho trình duyệt biết 
// là API Gateway này chấp nhận lời gọi từ các App/Domain độc lập (Ví dụ: React App)
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")

		// Trình duyệt thường gửi method OPTIONS trước để chọc dò, nên mình đồng ý phản hồi luôn
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
