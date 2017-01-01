package gonb

import (
	//"errors"
	"log"
	"testing"
)

func TestPeopleIndex(t *testing.T) {

	log.Println("Testing people indexing")

	var err error
	var params map[string]string
	params = make(map[string]string)
	params["limit"] = "2"

	a := NewAPI("risepartydev", NB_OAUTH_ACCESS_TOKEN)

	p := new(People)

	if err = p.Get(a, &params); err != nil {
		t.Error(err.Error())
	}

	if err = p.GetNext(a); err != nil {
		t.Error(err.Error())
	}
}
