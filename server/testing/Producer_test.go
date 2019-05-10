package testing

import (
	"AMCO/server/core/services"
	"log"
	"strings"
	"testing"
)

func TestTopicHasEndFix(t *testing.T) {
	if !services.TopicHasEndFix("test-sejfk32-234-req") {
		t.Error("Wrong!!, you should return True when topic has endfix -req")
	}
}

func TestFindTopicBody(t *testing.T) {
	const proposition = "geow-werkejf-wefwe-fwe-req-ewf-ef-wefwef-req"
	const objective = "geow-werkejf-wefwe-fwe-req-ewf-ef-wefwef"
	shatters := services.FindTopicBody(proposition)
	if !strings.Contains(shatters, objective) {
		t.Errorf("Wrong!!, your regexp can't find  %s in %s", objective, proposition)
	}

}

func TestMidgTopicNames(t *testing.T) {
	const proposition = "geow-werkejf-wefwe-fwe-req-ewf-ef-wefwef-req"
	const objective = "geow-werkejf-wefwe-fwe-req-ewf-ef-wefwef"
	shatter := services.MidgTopicNames(proposition)
	if shatter != objective {
		t.Errorf("wrong!!, your regexp or logic is bad you must return %s", objective)
	}
}

func TestA(t *testing.T) {
	a := []byte(`{\"result\":{\"message\":{\"name\":\"nuke\"},\"clientID\":\"3b5c8168-993b-4960-9ce5-a2b16fb3cc36\"},\"error\":\"\",\"statusCode\":200}`)
	log.Fatal(string	(a))
}
