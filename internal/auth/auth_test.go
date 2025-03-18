package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	header1 := make(http.Header)
	header1.Set("Authorization", "ApiKey 1234testkey5678")
	header2 := make(http.Header)
	header2.Set("Authorization", "FalseKey 13235553232")
	header3 := make(http.Header)

	tests := []struct {
		name    string
		header  http.Header
		want    string
		wantErr bool
	}{
		{
			name:    "Correct header",
			header:  header1,
			want:    "1234testkey5678",
			wantErr: false,
		},
		{
			name:    "Incorrect header",
			header:  header2,
			want:    "",
			wantErr: true,
		},
		{
			name:    "Empty header",
			header:  header3,
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.header)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() got = %v, expected %v", got, tt.want)
			}
		})
	}
}
