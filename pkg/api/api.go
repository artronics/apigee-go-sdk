package api

import (
	"fmt"
	"github.com/artronics/apigee/pkg"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

var httpClient = http.Client{}

type ApigeeApiConfig struct {
	pkg.ApigeeConfig
	Name string
}

func Get(config pkg.ApigeeConfig, resourceType pkg.ApigeeResource, resource interface{}) (io.ReadCloser, error) {
	var apps io.ReadCloser

	switch resourceType {
	case pkg.ApigeeApi:
		data := resource.(pkg.Api)

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
