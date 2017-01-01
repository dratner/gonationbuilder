package gonb

import (
	//"errors"
	"log"
	"net/url"
	"testing"
)

func TestShow(t *testing.T) {

	log.Println("Testing person show")
	/*
		p := new(Person)
		p.Id = 1
		err := p.Show()

		if err != nil {
			t.Error(err.Error())
		}
	*/
	u := "/api/v1/people?__nonce=sZDjoLaIT7jv9e08Mbmd9Q&__token=ANS_vP7mlocWfQ_ZI4s30fPW52u5AtrdZLSIqzBIPhay&limit=1"

	var res url.Values
	var err error

	if res, err = url.ParseQuery(u); err != nil {
		t.Error(err.Error())
	}

	log.Printf("Vals: %v\n", res)
}
