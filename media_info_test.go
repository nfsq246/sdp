package sdp

import (
	"fmt"
	"testing"
)

func TestMediaInfo(t *testing.T) {
	tests := []struct {
		item    string
		wantErr bool
	}{
		{"audio 19998 RTP/AVP 0 101", false},
		{"video 0 RTP/AVP 96", false},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			mi, err := parseMediaInfo(tt.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("MediaInfo error = %v, wantErr %v", err, tt.wantErr)
			}
			if str := mi.String(); str != tt.item {
				t.Errorf("MediaInfo string = %v, wantString %v", str, tt.item)
			}
		})
	}
}
