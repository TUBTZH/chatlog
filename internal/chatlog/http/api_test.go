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

// TestChatLogAPI_ChatLog 聊天记录查询API测试
func TestChatLogAPI_ChatLog(t *testing.T) {
	tests := []struct {
		name       string
		query      string
		wantStatus int
		wantKeys   []string
	}{
		{
			name:       "query with valid date range",
			query:      "time=2025-01-01~2026-12-31&talker=19631256769@chatroom&limit=5",
			wantStatus: http.StatusOK,
		},
		{
			name:       "query without time param",
			query:      "talker=test",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "query without talker param",
			query:      "time=2025-01-01",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "query with last-7d",
			query:      "time=last-7d&talker=test&limit=10",
			wantStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := http.Get(testServerAddr + "/api/v1/chatlog?" + tt.query)
			if err != nil {
				t.Skipf("Server not available: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tt.wantStatus {
				t.Errorf("Expected status %d, got %d", tt.wantStatus, resp.StatusCode)
			}
		})
	}
}
