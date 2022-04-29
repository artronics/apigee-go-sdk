package api

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func Create(resourceType ApigeeResource, resource interface{}) (io.ReadCloser, error) {
	switch resourceType {
	case Api:
		data := resource.(ApiData)
		multipartHeader, body, err := createForm("bundle", data.ZipBundle)
		if err != nil {
			return nil, err
		}
		req, _ := http.NewRequest("POST", data.url(), body)

		q := url.Values{}
		q.Add("name", data.Name)
		q.Add("action", data.Action)
		req.URL.RawQuery = q.Encode()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", data.Token))
		req.Header.Set("Content-Type", multipartHeader)

		res, err := httpClient.Do(req)
		if err != nil {
			return nil, err
		}
		if res.StatusCode != 201 {
			return nil, fmt.Errorf("GET request failed - %s", res.Status)
		}

		return res.Body, nil

	default:
		panic("unsupported Apigee resource type")
	}
}
