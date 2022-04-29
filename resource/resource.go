package resource

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

var httpClient = http.Client{}

type ApigeeResource int

const (
	_ ApigeeResource = iota
	Api
)

type operation int

const (
	_ = iota
	get
	list
	create
)

type Apigee struct {
	Token   string
	BaseUrl string
}

type Organization struct {
	Apigee
	Name string
}

func (o *Organization) url() string {
	return fmt.Sprintf("%s/organizations/%s", o.BaseUrl, o.Name)
}

type ApiData struct {
	Organization
	Name             string
	IncludeRevisions bool
	IncludeMetaData  bool
	ZipBundle        string
	Action           string
}

func (a *ApiData) url() string {
	return fmt.Sprintf("%s/apis", a.Organization.url())
}

func (a *ApiData) request(opt operation) (req *http.Request, err error) {
	path := a.url()
	defer func() {
		if req != nil {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.Token))
		}
	}()

	switch opt {
	case get:
		path = fmt.Sprintf("%s/%s", path, a.Name)

		return http.NewRequest("GET", path, nil)

	case list:
		req, err = http.NewRequest("GET", path, nil)
		if err != nil {
			return nil, err
		}

		q := url.Values{}
		q.Add("includeMetaData", strconv.FormatBool(a.IncludeMetaData))
		q.Add("includeRevisions", strconv.FormatBool(a.IncludeRevisions))
		req.URL.RawQuery = q.Encode()

		return req, nil

	case create:
		multipartHeader, body, err := createForm("bundle", a.ZipBundle)
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequest("POST", path, body)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Content-Type", multipartHeader)

		q := url.Values{}
		q.Add("name", a.Name)
		q.Add("action", a.Action)
		req.URL.RawQuery = q.Encode()

		return req, nil

	default:
		panic("unreachable code")
	}
}
