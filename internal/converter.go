package internal

import (
	"encoding/base64"
	"fmt"
	"os"
	"net/url"
	"strings"
)

func Converter(httpConfig HttpRequestConfig) CurlInput {

	var httpMethod string

	if httpConfig.Method == "" {
		httpMethod = "GET"
	} else {
		httpMethod = strings.ToUpper(httpConfig.Method)
	}

	queryParams := generateQueryParamsString(httpConfig.QueryParams)

	urlBase := substitutePathParams(httpConfig.Url, httpConfig.PathParams)

	httpUrl := urlBase + queryParams

	var basicAuthHeader string

	if len(httpConfig.BasicAuth.Username) > 0 {
		basicAuthHeader = generateBasicAuthHeader(httpConfig.BasicAuth)
	} else {
		basicAuthHeader = ""
	}

	var headers map[string]string

	if httpConfig.Headers == nil {
		headers = make(map[string]string)
	} else {
		headers = httpConfig.Headers
	}

	if basicAuthHeader != "" {
		headers["Authorization"] = basicAuthHeader
	}

	var payload string
	if httpConfig.PayloadPath != "" {

		fileBytes, err := os.ReadFile(httpConfig.PayloadPath)

		if err != nil {
			panic(err.Error())
		}

		payload = string(fileBytes)

	} else {
		payload = ""
	}

	return CurlInput{
		Url:        httpUrl,
		HttpMethod: httpMethod,
		Headers:    headers,
		Payload:    payload,
	}

}

func substitutePathParams(urlInput string, pathParams map[string]string) string {

	u, err := url.Parse(urlInput)
	if err != nil {
		panic(err.Error())
	}

	var preSubstitutePathParts []string = strings.Split(u.Path, "/")

	var postSubstitutePathParts []string = make([]string, 0)

	for _, pathPart := range preSubstitutePathParts {

		finalPathPart := substituteOrDefaultPathPart(pathPart, pathParams)
		postSubstitutePathParts = append(postSubstitutePathParts, finalPathPart)

	}

	finalPathParts := strings.Join(postSubstitutePathParts, "/")

	return fmt.Sprintf("%v://%v%v", u.Scheme, u.Host, finalPathParts)
}

func substituteOrDefaultPathPart(pathPart string, pathParams map[string]string) string {

	for placeholderWithoutPrefix, value := range pathParams {

		placeholder := ":" + placeholderWithoutPrefix

		if pathPart == placeholder {
			return value
		}

	}

	return pathPart

}

func generateQueryParamsString(queryParams map[string][]string) string {

	if len(queryParams) == 0 {
		return ""
	}

	var paramList []string

	for k, v := range queryParams {

		for _, value := range v {

			paramString := fmt.Sprintf("%v=%v", k, value)

			paramList = append(paramList, paramString)

		}

	}

	return fmt.Sprintf("?%v", strings.Join(paramList, "&"))
}

func generateBasicAuthHeader(basicAuth BasicAuth) string {
	rawCredentials := fmt.Sprintf("%v:%v", basicAuth.Username, basicAuth.Password)

	data := []byte(rawCredentials)

	dst := make([]byte, base64.StdEncoding.EncodedLen(len(data)))

	base64.StdEncoding.Encode(dst, data)

	return fmt.Sprintf("Basic %v", string(dst))
}
