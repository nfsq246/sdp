package sdp

import (
	"fmt"
	"testing"
)

func TestUtilParse(t *testing.T) {
	tests := []struct {
		item      string
		wantKey   string
		wantValue string
		wantErr   bool
	}{
		{"v=0", "v", "0", false},
		{"a=control:*", "a", "control:*", false},
		{"b=RS=600", "b", "RS=600", false},
		{"c", "", "", true},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			gotKey, gotValue, err := parse(tt.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotKey != tt.wantKey {
				t.Errorf("parse() gotKey = %v, want %v", gotKey, tt.wantKey)
			}
			if gotValue != tt.wantValue {
				t.Errorf("parse() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}
