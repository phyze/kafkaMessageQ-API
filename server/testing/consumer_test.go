// +build !prod

package testing

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func TestConsumeTopicThatDoesNotExists(t *testing.T) {

	url := "http://localhost:8000/amco/api/consumer"

	data := []byte(`{
		"autoOffsetReset":"earliest",
		"topics":["test-res"],
		"groupID":"a"
	}`)

	contentType := "application/josn"

	res, err := http.Post(url, contentType, bytes.NewBuffer(data))
	if err != nil {
		t.Error(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err.Error())
	}
	x := make(map[string]interface{})
	err = jsoniter.Unmarshal(body, &x)
	if err != nil {
		t.Error(err)
	}
	if x["error"] == "Broker: Unknown topic or partition" {
		t.Error(x["error"])
	}

}
