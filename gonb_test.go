package gonb

import (
	"log"
	"testing"
)

func TestAPI(t *testing.T) {
	log.Println("Testing API")

	var err error

	p := new(People)
	params := make(map[string]string)
	params["limit"] = "1"

	a := NewAPI("risepartydev", NB_OAUTH_ACCESS_TOKEN)
	if err = a.Exec("people", p, "GET", &params); err != nil {
		t.Error(err.Error())
	}
}
