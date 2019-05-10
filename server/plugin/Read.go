package plugin

import (
	"log"
	"os"
)

func ReadENV(envs *[]string) map[string]string {
	obj := make(map[string]string)
	for _, env := range *envs {
		obj[env] = os.Getenv(env)
	}
	log.Fatal(obj)
	return obj
}
