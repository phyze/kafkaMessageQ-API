// +build !prod

package testing

import (
	"strings"
	"testing"

	"KafkaMessageQ-API/server/plugin"
)

func TestStringFormat(t *testing.T) {
	result := plugin.StringFormat(
		"{{.hello}}",
		map[string]interface{}{
			"hello": "google",
		},
	)

	if strings.Compare(result, "google") == -1 {
		t.Error("message formated failure", result)
	}

}
