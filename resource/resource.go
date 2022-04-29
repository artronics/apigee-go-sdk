package resource

import (
	"fmt"
	"net/http"
)

var httpClient = http.Client{}

type ApigeeResource int

const (
	_ ApigeeResource = iota
	Proxy
	App
)

type operation int

const (
	_ operation = iota
	get
	list
	create
	// update
	deleteOpt
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
