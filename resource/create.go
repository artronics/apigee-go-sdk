package resource

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func Create(resourceType ApigeeResource, resource interface{}) (body io.ReadCloser, err error) {
	var req *http.Request

	switch resourceType {
	case Api:
		data := resource.(ApiData)
		req, err = data.request(create)

	default:
		panic("unsupported/wrong Apigee resource type")
	}

	if err != nil {
		return nil, err
	}

	var res *http.Response
	res, err = httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 201 {
		return nil, fmt.Errorf("create resource failed - %s", res.Status)
	}

	return res.Body, nil
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
