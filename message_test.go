package sdp

import (
	"fmt"
	"strings"
	"testing"
)

func TestMessage(t *testing.T) {
	tests := []struct {
		item    string
		wantErr bool
	}{
		{`v=0
		o=1011 134535 134535 IN IP4 192.168.0.100
		s=VaxSoft
		c=IN IP4 192.168.0.100
		t=0 0
		m=audio 25744 RTP/AVP 0 8 18 3 97 101
		a=rtpmap:0 PCMU/8000
		a=rtpmap:8 PCMA/8000
		a=rtpmap:18 G729/8000
		a=rtpmap:3 GSM/8000
		a=rtpmap:97 iLBC/8000
		a=fmtp:18 annexb=no
		a=rtpmap:101 telephone-event/8000
		a=fmtp:101 0-16
		a=sendrecv
		m=video 24180 RTP/AVP 96
		a=rtpmap:96 VP8/90000
		a=sendrecv
		`, false},
		{`v=0
		o=1010 437076426 8201 IN IP4 192.168.0.102
		s=SIP Implement by NTIL@NTUT
		c=IN IP4 192.168.0.102
		t=0 0
		m=audio 19998 RTP/AVP 0 101
		c=IN IP4 192.168.118.1
		a=rtpmap:0 PCMU/8000
		a=rtpmap:101 telephone-event/8000
		a=fmtp:101 0-15
		m=video 0 RTP/AVP 96
		a=sendrecv
		`, false},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			content := strings.ReplaceAll(tt.item, "\t", "")
			content = strings.ReplaceAll(content, "\n", CRLF)
			m, err := NewMessage(content)
			if (err != nil) != tt.wantErr {
				t.Errorf("Message error = %v, wantErr %v", err, tt.wantErr)
			}
			if str := m.String(); str != content {
				t.Errorf("Message string = %v, wantString %v", str, content)
			}
		})
	}
}
