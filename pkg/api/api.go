package api

import (
	"fmt"
	"github.com/artronics/apigee/pkg"
	"io"
	"log"
	"net/http"
)

var httpClient = http.Client{}

func List(config pkg.ApigeeConfig, resource string) (io.ReadCloser, error) {
	var apps io.ReadCloser

	if resource == "api" {
		baseUrl := fmt.Sprintf("%s/organizations/%s/apis", config.BaseUrl, config.Organization)
		log.Println(baseUrl)

		req, err := http.NewRequest("GET", baseUrl, nil)
		if err != nil {
			return apps, err
		}

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.Token))

		res, err := httpClient.Do(req)
		if err != nil {
			return apps, err
		}
		if res.StatusCode != 200 {
			return apps, fmt.Errorf("GET request failed with status code %d", res.StatusCode)
		}
		return res.Body, nil

		/*		body, err := ioutil.ReadAll(res.Body)
				if err != nil {
					return apps, err
				}

				err = json.Unmarshal(body, &apps)
				if err != nil {
					return apps, err
				}

				return apps, nil
		*/
	}

	return apps, nil
}

type ApigeeApp struct {
	pkg.ApigeeConfig
	organization string
}

/*
func NewApigeeApp(config ApigeeConfig) (*ApigeeApp, error) {
	if config.organization == "" {
		return nil, errors.New("organization is required, but it was not provided")
	}
	baseUrl := fmt.Sprintf("%s/organizations/%s/apihs", config.baseUrl, config.organization)
	return &ApigeeApp{
		C:  config,
		baseUrl: baseUrl,
	}, nil
}
*/
/*
func (a ApigeeApp) list() (Apps, error) {
	var apps Apps

	req, err := http.NewRequest("GET", a.baseUrl, nil)
	if err != nil {
		return apps, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.C.token))

	res, err := httpClient.Do(req)
	if err != nil {
		return apps, err
	}
	if res.StatusCode != 200 {
		return apps, fmt.Errorf("GET request failed with status code %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return apps, err
	}

	err = json.Unmarshal(body, &apps)
	if err != nil {
		return apps, err
	}

	return apps, nil
}*/
