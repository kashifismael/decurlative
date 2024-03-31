package internal

import (
	"bytes"
	"fmt"
	"strings"
)

func Converter(httpConfig HttpRequestConfig) CurlInput {

	var httpMethod string

	if httpConfig.Method == "" {
		httpMethod = "GET"
	} else {
		httpMethod = strings.ToUpper(httpConfig.Method)
	}

	queryParams := generateQueryParamsString(httpConfig)

	urlBase := substitutePathParams(httpConfig.Url, httpConfig.PathParams)

	httpUrl := urlBase + queryParams

	return CurlInput{
		Url:        httpUrl,
		HttpMethod: httpMethod,
	}

}

func substitutePathParams(url string, pathParams map[string]string) string {

	var urlPointer = new(string)
	*urlPointer = url

	for k, v := range pathParams {

		urlValue := fmt.Sprint(*urlPointer)

		placeholder := fmt.Sprintf(":%v", k)

		*urlPointer = strings.ReplaceAll(urlValue, placeholder, v)
	}

	urlValue := *urlPointer

	return fmt.Sprint(urlValue)
}

func generateQueryParamsString(httpConfig HttpRequestConfig) string {
	var queryParamsBytes bytes.Buffer

	queryParamsBytes.WriteString("?")

	for k, v := range httpConfig.QueryParams {

		for _, value := range v {

			paramString := fmt.Sprintf("%v=%v&", k, value)

			queryParamsBytes.WriteString(paramString)

		}

	}

	return queryParamsBytes.String()
}
