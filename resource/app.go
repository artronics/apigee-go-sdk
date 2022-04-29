package resource

import (
	"fmt"
	"net/http"
)

type AppData struct {
	Organization
	Id string
}

func (a *AppData) url() string {
	return fmt.Sprintf("%s/apps", a.Organization.url())
}

func (a *AppData) request(opt operation) (req *http.Request, err error) {
	path := a.url()
	defer func() {
		if req != nil {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.Token))
		}
	}()

	switch opt {
	case get:
		path = fmt.Sprintf("%s/%s", path, a.Id)

		return http.NewRequest("GET", path, nil)

	case list:
		return http.NewRequest("GET", path, nil)

	default:
		panic("operation not supported")
	}
}
