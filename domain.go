package main

type HttpRequestConfig struct {
	Url         string              `json:"url"`
	Method      string              `json:"method"`
	QueryParams map[string][]string `json:"queryParams"`
}

type CurlInput struct {
	Url        string
	HttpMethod string
}
