package resource

import (
	"fmt"
	"io"
	"net/http"
)

func List(resourceType ApigeeResource, resource interface{}) (body io.ReadCloser, err error) {
	var req *http.Request

	switch resourceType {
	case Proxy:
		data := resource.(ProxyData)
		req, err = data.request(list)

	case App:
		data := resource.(AppData)
		req, err = data.request(list)

	default:
		panic("unsupported/wrong Apigee resource type")
	}

	if err != nil {
		return nil, err
	}

	var res *http.Response
	res, err = httpClient.Do(req)
	if err != nil {
		return body, err
	}

	if res.StatusCode != 200 {
		return body, fmt.Errorf("list resource failed - %s", res.Status)
	}

	return res.Body, nil
}
