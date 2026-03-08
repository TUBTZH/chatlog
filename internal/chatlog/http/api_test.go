package http

import (
	"net/http"
	"testing"
)

const (
	testServerAddr = "http://127.0.0.1:5030"
)

// APIResponse 通用API响应结构
type APIResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// TestHelper 测试辅助函数
type TestHelper struct {
	serverURL string
}

// NewTestHelper 创建测试辅助函数
func NewTestHelper(serverURL string) *TestHelper {
	return &TestHelper{serverURL: serverURL}
}

// Get 发送GET请求
func (h *TestHelper) Get(endpoint string) (*http.Response, error) {
	return http.Get(h.serverURL + endpoint)
}

// TestChatLogAPI_Health 健康检查测试
func TestChatLogAPI_Health(t *testing.T) {
	// 测试服务器是否可访问
	resp, err := http.Get(testServerAddr + "/health")
	if err != nil {
		t.Skipf("Server not available: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}
