package gonb

import (
	//"errors"
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
	//	if err = a.Exec("people", p, "GET", params); err != nil {
	//		t.Error(err.Error())
	//	}
	log.Printf("RES %v\n", p)
	log.Println(p.Next)
	return
	if err = a.Exec(p.Next, p, "GET", params); err != nil {
		t.Error(err.Error())
	}

}
