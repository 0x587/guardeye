package httpcb

import "testing"

func Test_topicMatch(t *testing.T) {
	tests := []struct {
		pattern string
		topic   string
		expect  bool
	}{
		{"/sensor/+/temperature", "/sensor/livingroom/temperature", true},
		{"/sensor/+/temperature", "/sensor/livingroom/humidity", false},
		{"/device/#", "/device/abc/status", true},
		{"/device/#", "/device", true},
		{"/device/#", "/dev", false},
		{"/sensor/+/+", "/sensor/a/b", true},
		{"/sensor/+/+", "/sensor/a/b/c", false},
	}
	for _, tt := range tests {
		t.Run(tt.pattern, func(t *testing.T) {
			if got := topicMatch(tt.pattern, tt.topic); got != tt.expect {
				t.Errorf("topicMatch() = %v, want %v", got, tt.expect)
			}
		})
	}
}
