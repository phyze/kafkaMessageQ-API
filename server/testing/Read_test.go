// +build !prod

package testing

import (
	"AMCO/server/plugin"
	"testing"
)

func TestReadEnv(t *testing.T) {
	s := []string{"AMCO_VERSION", "AMCO_APPNAME"}
	m := plugin.ReadENV(&s)

	if m["AMCO_APPNAME"] == "" {
		t.Error(" this function work failed because AMCO_APPNAME is empty")
	}
	if m["AMCO_VERSION"] == "" {
		t.Error(" this function work failed because AMCO_VERSION is empty")
	}

}
