// +build !prod

package testing

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestTimeout(t *testing.T) {
	url := "http://localhost:8000/amco/api/producer"
	// data := []byte(`{
	// 	"aWait":false,
	// 	"message":{"mylove":"nuke"},
	// 	"topics":{"sending":"test-req"}
	// }`)
	data := []byte(`{
		"aWait":true,
		"message":{"mylove":"nuke"},
		"topics":{"sending":"test-req","receiving":"test-res"},
	"timeout":30
	}`)

	contentType := "application/josn"

	res, err := http.Post(url, contentType, bytes.NewBuffer(data))
	if err != nil {
		log.Fatalln(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Fatal(string(body))
}
