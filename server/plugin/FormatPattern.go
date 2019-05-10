package plugin

import (
	"bytes"
	"text/template"
)

func StringFormat(fmt string, args map[string]interface{}) (str string) {
	var msg bytes.Buffer
	tmpl, err := template.New("errmsg").Parse(fmt)
	if err != nil {
		return fmt
	}
	tmpl.Execute(&msg, args)
	return msg.String()
}

func ErrorResponseJson(err error) string {
	return `{"error":"` + err.Error() + `"}`
}
