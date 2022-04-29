package resource

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type ProxyData struct {
	Organization
	Name             string
	IncludeRevisions bool
	IncludeMetaData  bool
	ZipBundle        string
	Action           string
}

func (p *ProxyData) url() string {
	return fmt.Sprintf("%s/apis", p.Organization.url())
}

func (p *ProxyData) request(opt operation) (req *http.Request, err error) {
	path := p.url()
	defer func() {
		if req != nil {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", p.Token))
		}
	}()

	switch opt {
	case get:
		path = fmt.Sprintf("%s/%s", path, p.Name)

		return http.NewRequest("GET", path, nil)

	case list:
		req, err = http.NewRequest("GET", path, nil)
		if err != nil {
			return nil, err
		}

		q := url.Values{}
		q.Add("includeMetaData", strconv.FormatBool(p.IncludeMetaData))
		q.Add("includeRevisions", strconv.FormatBool(p.IncludeRevisions))
		req.URL.RawQuery = q.Encode()

		return req, nil

	case create:
		multipartHeader, body, err := createForm("bundle", p.ZipBundle)
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequest("POST", path, body)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Content-Type", multipartHeader)

		q := url.Values{}
		q.Add("name", p.Name)
		q.Add("action", p.Action)
		req.URL.RawQuery = q.Encode()

		return req, nil

	case deleteOpt:
		path = fmt.Sprintf("%s/%s", path, p.Name)

		return http.NewRequest("DELETE", path, nil)
	default:
		panic("operation not supported")
	}
}
