package main

import (
	"log"

	"github.com/dinhbaokhanh/Final-Project-API-Gateway/internal/app"
	"github.com/dinhbaokhanh/Final-Project-API-Gateway/internal/config"
)

func main() {
	// Bước 1: Đọc và giải mã rễ file cấu hình JSON của toàn bộ Gateway
	cfg, err := config.Load("gateway.json")
	if err != nil {
		log.Fatalf("Lỗi không thể tải file cấu hình Gateway: %v", err)
	}

	// Bước 2: Khởi tạo ứng dụng Gateway (Core App) dựa trên cấu hình đã đọc
	gateway, err := app.New(cfg)
	if err != nil {
		log.Fatalf("Lỗi trong quá trình kết nối Core Gateway: %v", err)
	}

	// Bước 3: Chạy Gateway và bắt đầu lắng nghe các Request từ tầng Frontend
	log.Printf("PTIT Gateway đã khởi động thành công và đang lắng nghe trên cổng :%d", cfg.Port)
	if err := gateway.Run(); err != nil {
		log.Fatalf("Gateway đã dừng hoạt động do xảy ra lỗi: %v", err)
	}
}
