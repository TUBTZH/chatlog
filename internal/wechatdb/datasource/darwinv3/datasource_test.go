package darwinv3

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fsnotify/fsnotify"
	"github.com/sjzar/chatlog/internal/wechatdb/datasource/dbm"
)

// mockDBM implements dbm.DBManagerInterface for testing
type mockDBM struct {
	db *sql.DB
}

func (m *mockDBM) GetDB(name string) (*sql.DB, error) {
	return m.db, nil
}

func (m *mockDBM) GetDBs(name string) ([]*sql.DB, error) {
	return []*sql.DB{m.db}, nil
}

func (m *mockDBM) GetDBPath(name string) ([]string, error) {
	return []string{}, nil
}

func (m *mockDBM) AddGroup(group *dbm.Group) error {
	return nil
}

func (m *mockDBM) OpenDB(path string) (*sql.DB, error) {
	return m.db, nil
}

func (m *mockDBM) AddCallback(group string, callback func(event fsnotify.Event) error) error {
	return nil
}

func (m *mockDBM) Start() error {
	return nil
}

func (m *mockDBM) Stop() error {
	return nil
}

func (m *mockDBM) Close() error {
	return nil
}

var _ dbm.DBManagerInterface = (*mockDBM)(nil)

func TestGetContacts_FuzzySearch(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock: %v", err)
	}
	defer db.Close()

	// 期望模糊查询
	mock.ExpectQuery("SELECT .+ FROM WCContact").
		WithArgs("%test%", "%test%", "%test%", "%test%").
		WillReturnRows(sqlmock.NewRows([]string{"m_nsUsrName", "nickname", "m_nsRemark", "m_uiSex", "m_nsAliasName"}).
			AddRow("testuser@weixin", "testnick", "testremark", 1, "testalias"))

	ds := &DataSource{dbm: &mockDBM{db: db}}
	ctx := context.Background()

	contacts, err := ds.GetContacts(ctx, "test", 0, 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(contacts) != 1 {
		t.Fatalf("expected 1 contact, got %d", len(contacts))
	}

	if contacts[0].UserName != "testuser@weixin" {
		t.Fatalf("expected username 'testuser@weixin', got '%s'", contacts[0].UserName)
	}
}

func TestGetChatRooms_FuzzySearch(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock: %v", err)
	}
	defer db.Close()

	// 期望模糊查询并关联contact表获取群名称
	mock.ExpectQuery("SELECT g\\..+ FROM GroupContact g LEFT JOIN WCContact co").
		WithArgs("%test%", "%test%", "%test%").
		WillReturnRows(sqlmock.NewRows([]string{"m_nsUsrName", "nickname", "m_nsRemark", "m_nsChatRoomMemList", "m_nsChatRoomAdminList", "co_nickname", "co_remark"}).
			AddRow("123456@chatroom", "测试群", "群备注", "", "", "测试群", "群备注"))

	ds := &DataSource{dbm: &mockDBM{db: db}}
	ctx := context.Background()

	chatRooms, err := ds.GetChatRooms(ctx, "test", 0, 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(chatRooms) != 1 {
		t.Fatalf("expected 1 chatroom, got %d", len(chatRooms))
	}

	if chatRooms[0].NickName != "测试群" {
		t.Fatalf("expected NickName '测试群', got '%s'", chatRooms[0].NickName)
	}
}
