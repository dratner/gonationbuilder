package gonb

import (
	"encoding/json"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"io"
	"net/http"
)

func NewAPI(slug string, token string) *API {
	a := new(API)
	a.Init(slug, token)
	return a
}

type API struct {
	access_token string
	slug         string
	cli          *http.Client
	ctx          context.Context
}

func (a *API) Init(slug string, token string) {
	a.access_token = token
	a.slug = slug
	a.ctx = *new(context.Context)
	a.cli = oauth2.NewClient(a.ctx, a)
}

func (a API) Token() (*oauth2.Token, error) {
	tkn := new(oauth2.Token)
	tkn.AccessToken = a.access_token
	return tkn, nil
}

//	Endpoint describes just the final part of the url (e.g. "people")

func (a *API) Exec(endpoint string, res interface{}, method string, params map[string]string) error {

	var body io.Reader
	var err error
	var resp *http.Response
	var req *http.Request
	var key, val, url string

	if endpoint != "" {
		url = "https://" + a.slug + ".nationbuilder.com/api/v1/" + endpoint + "?access_token=" + a.access_token

		for key, val = range params {
			url += "&" + key + "=" + val
		}
	}

	if req, err = http.NewRequest(method, url, body); err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	if resp, err = a.cli.Do(req); err != nil {
		return err
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	if err = decoder.Decode(res); err != nil {
		return err
	}

	return nil
}
