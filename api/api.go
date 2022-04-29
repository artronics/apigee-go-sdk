package api

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

var httpClient = http.Client{}

func Get(config ApigeeConfig, resourceType ApigeeResource, resource interface{}) (io.ReadCloser, error) {
	var apps io.ReadCloser

	switch resourceType {
	case Api:
		data := resource.(ApiData)

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
}

func Create(config ApigeeConfig, resourceType ApigeeResource, resource interface{}) (io.ReadCloser, error) {
	switch resourceType {
	case Api:
		data := resource.(ApiData)
		multipartHeader, body, err := createForm("bundle", data.ZipBundle)
		if err != nil {
			return nil, err
		}
		baseUrl := fmt.Sprintf("%s/organizations/%s/apis", config.BaseUrl, data.Organization.Name)
		req, _ := http.NewRequest("POST", baseUrl, body)

		q := url.Values{}
		q.Add("name", data.Name)
		q.Add("action", data.Action)
		req.URL.RawQuery = q.Encode()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.Token))
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

func createForm(key string, val string) (_ string, _ io.Reader, err error) {
	body := new(bytes.Buffer)
	mp := multipart.NewWriter(body)
	defer func(mp *multipart.Writer) {
		err = mp.Close()
	}(mp)

	file, err := os.Open(val)
	if err != nil {
		return "", nil, err
	}
	defer func(file *os.File) {
		err = file.Close()
	}(file)

	part, err := mp.CreateFormFile(key, val)
	if err != nil {
		return "", nil, err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return "", nil, err
	}

	return mp.FormDataContentType(), body, err
}
