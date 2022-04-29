package api

type ApigeeConfig struct {
	Token   string
	BaseUrl string
}

type ApigeeResource int

const (
	_ ApigeeResource = iota
	Api
)

type Organization struct {
	Name string
}

type ApiData struct {
	Organization
	Name             string
	IncludeRevisions bool
	IncludeMetaData  bool
	ZipBundle        string
	Action           string
}
