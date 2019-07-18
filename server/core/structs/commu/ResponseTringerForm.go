package commu

import (
	jsoniter "github.com/json-iterator/go"
)

type ResponseTringger struct {
	Result     interface{} `json:"result"`
	Error      string      `json:"error"`
	StatusCode int         `json:"statusCode"`
}

func (rt *ResponseTringger) ToJson() []byte {
	data, _ := jsoniter.Marshal(rt)

	return data
}
