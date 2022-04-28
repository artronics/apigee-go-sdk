package api

type ApigeeConfig struct {
	Token   string
	BaseUrl string
}

type ApigeeResource int

const (
	_ ApigeeResource = iota
	ApigeeApi
)

type Organization struct {
	Name string
}

type Api struct {
	Organization
	Name             string
	IncludeRevisions bool
	IncludeMetaData  bool
}
