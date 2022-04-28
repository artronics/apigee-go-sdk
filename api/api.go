package api

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

var httpClient = http.Client{}

func Get(config ApigeeConfig, resourceType ApigeeResource, resource interface{}) (io.ReadCloser, error) {
	var apps io.ReadCloser

	switch resourceType {
	case ApigeeApi:
		data := resource.(Api)

		baseUrl := fmt.Sprintf("%s/organizations/%s/apis/%s", config.BaseUrl, data.Organization.Name, data.Name)

		req, err := http.NewRequest("GET", baseUrl, nil)
		if err != nil {
			return apps, err
		}

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.Token))

		q := url.Values{}
		q.Add("includeMetaData", strconv.FormatBool(data.IncludeMetaData))
		q.Add("includeRevisions", strconv.FormatBool(data.IncludeRevisions))
		req.URL.RawQuery = q.Encode()

		res, err := httpClient.Do(req)
		if err != nil {
			return apps, err
		}
		if res.StatusCode != 200 {
			return apps, fmt.Errorf("GET request failed - %s", res.Status)
		}
		return res.Body, nil

	default:
		panic("unsupported Apigee resource type")
	}

	return apps, nil
}
