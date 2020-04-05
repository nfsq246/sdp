package sdp

import (
	"fmt"
	"testing"
)

func TestBasic(t *testing.T) {
	tests := []struct {
		item    string
		wantErr bool
	}{
		{"1011 134535 134535 IN IP4 192.168.0.100", false},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			item, err := parseBasic(tt.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("basic error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			wantString := output(FieldBasic, tt.item)
			if str := item.String(); str != wantString {
				t.Errorf("basic string = %v, wantString %v", str, wantString)
			}
		})
	}
}
