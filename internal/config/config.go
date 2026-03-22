package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// GatewayConfig là cấu trúc gốc chứa toàn bộ cài đặt của Gateway (đọc từ file JSON)
type GatewayConfig struct {
	Port      int              `json:"port"`
	Endpoints []EndpointConfig `json:"endpoints"`
}

// EndpointConfig định nghĩa một API Route mà Gateway sẽ mở ra để Frontend gọi
type EndpointConfig struct {
	Endpoint string          `json:"endpoint"`
	Method   string          `json:"method"`
	Backend  []BackendConfig `json:"backend"`
}

// BackendConfig chứa thông tin về các Service phía sau tương ứng với Endpoint phía trên
type BackendConfig struct {
	Host       []string `json:"host"`
	URLPattern string   `json:"url_pattern"`
}

// Load đọc và giải mã file cấu hình JSON của Gateway
func Load(configPath string) (*GatewayConfig, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("không thể mở file cấu hình: %w", err)
	}
	defer file.Close()

	var config GatewayConfig
	decoder := json.NewDecoder(file)
	
	// Giải mã nội dung JSON đắp vào struct GatewayConfig
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("gặp lỗi cú pháp khi parse cấu hình JSON: %w", err)
	}

	return &config, nil
}
