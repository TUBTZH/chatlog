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

// TestChatLogAPI_ContactList 联系人列表API测试
func TestChatLogAPI_ContactList(t *testing.T) {
	tests := []struct {
		name       string
		query      string
		wantStatus int
	}{
		{
			name:       "get all contacts",
			query:      "",
			wantStatus: http.StatusOK,
		},
		{
			name:       "search contacts with keyword",
			query:      "keyword=test",
			wantStatus: http.StatusOK,
		},
		{
			name:       "search with limit",
			query:      "limit=10",
			wantStatus: http.StatusOK,
		},
		{
			name:       "search with limit and offset",
			query:      "limit=10&offset=10",
			wantStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			endpoint := "/api/v1/contact"
			if tt.query != "" {
				endpoint += "?" + tt.query
			}
			resp, err := http.Get(testServerAddr + endpoint)
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

// TestChatLogAPI_ChatRoomList 群聊列表API测试
func TestChatLogAPI_ChatRoomList(t *testing.T) {
	tests := []struct {
		name       string
		query      string
		wantStatus int
	}{
		{
			name:       "get all chatrooms",
			query:      "",
			wantStatus: http.StatusOK,
		},
		{
			name:       "search chatrooms with keyword",
			query:      "keyword=4403",
			wantStatus: http.StatusOK,
		},
		{
			name:       "search with fuzzy match",
			query:      "keyword=康复",
			wantStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			endpoint := "/api/v1/chatroom"
			if tt.query != "" {
				endpoint += "?" + tt.query
			}
			resp, err := http.Get(testServerAddr + endpoint)
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

// TestChatLogAPI_SessionList 会话列表API测试
func TestChatLogAPI_SessionList(t *testing.T) {
	tests := []struct {
		name       string
		query      string
		wantStatus int
	}{
		{
			name:       "get all sessions",
			query:      "",
			wantStatus: http.StatusOK,
		},
		{
			name:       "search sessions with keyword",
			query:      "keyword=test",
			wantStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			endpoint := "/api/v1/session"
			if tt.query != "" {
				endpoint += "?" + tt.query
			}
			resp, err := http.Get(testServerAddr + endpoint)
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

// TestChatLogAPI_FuzzySearch 模糊搜索功能测试
func TestChatLogAPI_FuzzySearch(t *testing.T) {
	tests := []struct {
		name       string
		endpoint   string
		query      string
		wantStatus int
	}{
		{
			name:       "chatroom fuzzy search by nickname",
			endpoint:   "/api/v1/chatroom",
			query:      "keyword=4403",
			wantStatus: http.StatusOK,
		},
		{
			name:       "chatroom fuzzy search by remark",
			endpoint:   "/api/v1/chatroom",
			query:      "keyword=康复",
			wantStatus: http.StatusOK,
		},
		{
			name:       "contact fuzzy search by alias",
			endpoint:   "/api/v1/contact",
			query:      "keyword=test",
			wantStatus: http.StatusOK,
		},
		{
			name:       "contact fuzzy search by remark",
			endpoint:   "/api/v1/contact",
			query:      "keyword=备注",
			wantStatus: http.StatusOK,
		},
		{
			name:       "contact fuzzy search by nickname",
			endpoint:   "/api/v1/contact",
			query:      "keyword=昵称",
			wantStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := http.Get(testServerAddr + tt.endpoint + "?" + tt.query)
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
