package internal

type HttpRequestConfig struct {
	Url         string              `json:"url"`
	Method      string              `json:"method"`
	QueryParams map[string][]string `json:"queryParams"`
	PathParams  map[string]string   `json:"pathParams"`
}

type CurlInput struct {
	Url        string
	HttpMethod string
}

type CommandLineFlags struct {
	DebugMode bool
}
