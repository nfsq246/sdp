package sdp

import (
	"fmt"
	"testing"
)

func TestConnection(t *testing.T) {
	tests := []struct {
		item    string
		wantErr bool
	}{
		{"IN IP4 192.168.0.100", false},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			item, err := parseConnection(tt.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("connection error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if str := item.String(); str != tt.item {
				t.Errorf("connection string = %v, wantString %v", str, tt.item)
			}
		})
	}
}
