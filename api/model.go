package api

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

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
	path := fmt.Sprintf("%s/apis", a.Organization.url())
	defer func() {
		if req != nil {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.Token))
		}
	}()
	//headers := http.Header{}
	//headers.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))

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

	case create:
	default:
		panic("unreachable code")
	}

	return req, nil
}
