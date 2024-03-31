package internal

type HttpRequestConfig struct {
	Url         string              `json:"url"`
	Method      string              `json:"method"`
	QueryParams map[string][]string `json:"queryParams"`
	PathParams  map[string]string   `json:"pathParams"`
	Headers     map[string]string   `json:"headers"`
	BasicAuth   BasicAuth           `json:"basicAuth"`
}

type CurlInput struct {
	Url        string
	HttpMethod string
	Headers    map[string]string
}

type CommandLineFlags struct {
	DebugMode bool
}

type BasicAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
