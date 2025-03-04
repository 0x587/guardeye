package ros

import (
	"fmt"
	"testing"
)

func Test_TopicAndName(t *testing.T) {
	topics := []string{
		"/abc_edf/ghi_jkl/ghi_jkl",
	}
	for i, topic := range topics {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			name := Topic2Name(topic)
			t.Log(name)
			if Name2Topic(name) != topic {
				t.Errorf("%s => %s\n%s => %s", topic, name, name, topic)
			}
		})
	}
}
