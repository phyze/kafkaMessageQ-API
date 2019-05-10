package plugin

import (
	"io/ioutil"
	"net/http"
)

//ReadBody  read []byte from  *http.Request
func ReadBody(req *http.Request) ([]byte, error) {
	
	data, err := ioutil.ReadAll(req.Body)
	return data, err
}
