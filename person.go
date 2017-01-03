package gonb

import (
	"errors"
	"log"
	"net/url"
	"strconv"
)

type Address struct {
	Zip string `json:"zip"`
}

type Person struct {
	Id             int64    `json:"id"`
	FirstName      string   `json:"first_name"`
	LastName       string   `json:"last_name"`
	Email          string   `json:"email"`
	Tags           []string `json:"tags"`
	PrimaryAddress Address  `json:"primary_address"`
	Volunteer      bool     `json:"is_volunteer"`
	Username       string   `json:"username"`
}

func (p *Person) Show() error {
	nb_url := NB_API_BASE_URL + "/v1/people/" + strconv.FormatInt(p.Id, 10) + "?access_token=" + NB_OAUTH_ACCESS_TOKEN
	log.Println(nb_url)
	return nil
}

type People struct {
	Results []Person `json:"results"`
	Prev    string   `json:"prev"`
	Next    string   `json:"next"`
}

func (p *People) Get(a *API, params *map[string]string) error {

	var err error

	if err = a.Exec("people", p, "GET", params); err != nil {
		return err
	}

	return nil
}

func (p *People) GetNext(a *API) error {

	var err error
	var u *url.URL
	var params map[string]string

	if p.Next == "" {
		return errors.New("No data to generate next.")
	}

	if u, err = url.Parse(p.Next); err != nil {
		return err
	}

	res := u.Query()

	params = make(map[string]string)

	for k, v := range res {
		switch k {
		case "__nonce":
			params["__nonce"] = v[0]
			break
		case "__token":
			params["__token"] = v[0]
			break
		case "limit":
			params["limit"] = v[0]
			break
		}
	}

	if err = a.Exec("people", p, "GET", &params); err != nil {
		return err
	}

	return nil
}
