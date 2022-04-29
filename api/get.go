package api

import (
	"fmt"
	"io"
	"net/http"
)

func Get(resourceType ApigeeResource, resource interface{}) (body io.ReadCloser, err error) {
	var req *http.Request
	var res *http.Response

	switch resourceType {
	case Api:
		data := resource.(ApiData)

		req, err = data.request(get)
		if err != nil {
			return nil, err
		}

	default:
		panic("unsupported Apigee resource type")
	}

	res, err = httpClient.Do(req)
	if err != nil {
		return body, err
	}

	if res.StatusCode != 200 {
		return body, fmt.Errorf("GET request failed - %s", res.Status)
	}

	return res.Body, nil
}
