// +build !prod

package testing

import (
	"KafkaMessageQ-API/server/plugin/uuid"
	"log"
	"reflect"
	"strings"
	"testing"
)

func TestSliceToString(t *testing.T) {
	s := &[]string{"square-req", "__confluent.support.metrics", "__consumer_offsets", "square-res"}
	log.Fatal(strings.Join(*s, " "))

}

func TestParseUUID(t *testing.T) {
	id, _ := uuid.NewUUID4()
	a, _ := parseUUID(id.String())
	log.Fatal(reflect.TypeOf(a))
}

func parseUUID(s string) (uuid.UUID, error) {
	uuid, err := uuid.Parse(s)
	return uuid, err
}
