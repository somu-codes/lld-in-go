package store

import (
	"testing"
)

func TestInMemoryStore_Get(t *testing.T) {
	tests := []struct {
		name     string
		initData map[string]string
		args     string
		wantUrl  string
		wantErr  bool
	}{
		{
			name: "existing short code",
			initData: map[string]string{
				"abc123": "https://example.com",
			},
			args:    "abc123",
			wantUrl: "https://example.com",
			wantErr: false,
		},
		{
			name:     "non-existing short code",
			initData: map[string]string{},
			args:     "notfound",
			wantUrl:  "",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &InMemoryStore{
				data:        tt.initData,
				reverseData: map[string]string{},
			}
			gotUrl, err := s.Get(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotUrl != tt.wantUrl {
				t.Errorf("Get() gotUrl = %v, want %v", gotUrl, tt.wantUrl)
			}
		})
	}
}

func TestInMemoryStore_Save(t *testing.T) {
	tests := []struct {
		name           string
		existingData   map[string]string
		existingRevMap map[string]string
		inputURL       string
		wantSameCode   bool
	}{
		{
			name:           "new URL",
			existingData:   map[string]string{},
			existingRevMap: map[string]string{},
			inputURL:       "https://newsite.com",
			wantSameCode:   false,
		},
		{
			name: "idempotent URL",
			existingData: map[string]string{
				"code123": "https://example.com",
			},
			existingRevMap: map[string]string{
				"https://example.com": "code123",
			},
			inputURL:     "https://example.com",
			wantSameCode: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &InMemoryStore{
				data:        tt.existingData,
				reverseData: tt.existingRevMap,
			}
			gotCode, err := s.Save(tt.inputURL)
			if err != nil {
				t.Errorf("Save() error = %v", err)
				return
			}
			if tt.wantSameCode {
				if gotCode != tt.existingRevMap[tt.inputURL] {
					t.Errorf("Save() expected same code %v, got %v", tt.existingRevMap[tt.inputURL], gotCode)
				}
			} else {
				if gotCode == "" {
					t.Error("Save() returned empty code for new URL")
				}
				if savedUrl := s.data[gotCode]; savedUrl != tt.inputURL {
					t.Errorf("Save() did not store URL correctly: got %v, want %v", savedUrl, tt.inputURL)
				}
			}
		})
	}
}
