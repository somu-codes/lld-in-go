package service

import (
	"errors"
	"testing"
	"url-shortener/internal/store"
)

// Manual mock implementation
type mockStore struct {
	saveFunc func(url string) (string, error)
	getFunc  func(code string) (string, error)
}

func (m *mockStore) Save(url string) (string, error) {
	return m.saveFunc(url)
}

func (m *mockStore) Get(code string) (string, error) {
	return m.getFunc(code)
}

func TestURLService_Resolve(t *testing.T) {
	tests := []struct {
		name    string
		fields  store.URLStore
		args    string
		wantUrl string
		wantErr bool
	}{
		{
			name: "existing code",
			fields: &mockStore{
				getFunc: func(code string) (string, error) {
					if code == "abc123" {
						return "https://example.com", nil
					}
					return "", errors.New("not found")
				},
			},
			args:    "abc123",
			wantUrl: "https://example.com",
			wantErr: false,
		},
		{
			name: "non-existent code",
			fields: &mockStore{
				getFunc: func(code string) (string, error) {
					return "", errors.New("not found")
				},
			},
			args:    "notexist",
			wantUrl: "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &URLService{
				store: tt.fields,
			}
			gotUrl, err := s.Resolve(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Resolve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotUrl != tt.wantUrl {
				t.Errorf("Resolve() gotUrl = %v, want %v", gotUrl, tt.wantUrl)
			}
		})
	}
}

func TestURLService_Shorten(t *testing.T) {
	tests := []struct {
		name     string
		fields   store.URLStore
		args     string
		wantCode string
		wantErr  bool
	}{
		{
			name: "successful shorten",
			fields: &mockStore{
				saveFunc: func(url string) (string, error) {
					if url == "https://example.com" {
						return "abc123", nil
					}
					return "", errors.New("save failed")
				},
			},
			args:     "https://example.com",
			wantCode: "abc123",
			wantErr:  false,
		},
		{
			name: "save fails",
			fields: &mockStore{
				saveFunc: func(url string) (string, error) {
					return "", errors.New("failed to save")
				},
			},
			args:     "bad-url",
			wantCode: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &URLService{
				store: tt.fields,
			}
			gotCode, err := s.Shorten(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Shorten() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCode != tt.wantCode {
				t.Errorf("Shorten() gotCode = %v, want %v", gotCode, tt.wantCode)
			}
		})
	}
}
