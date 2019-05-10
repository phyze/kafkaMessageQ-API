// +build !prod

package testing

import (
	"AMCO/server/plugin"
	"log"
	"testing"
)

func TestCompare(t *testing.T) {
	_, err := plugin.Isin("b", []string{"a", "b"})
	if err != nil {
		log.Fatal(err.Error())
	}

}
