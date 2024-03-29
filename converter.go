package main

import (
	"bytes"
	"fmt"
	"strings"
)

func Converter(httpConfig HttpRequestConfig) CurlInput {

	httpMethod := strings.ToUpper(httpConfig.Method)

	queryParams := generateQueryParamsString(httpConfig)

	httpUrl := httpConfig.Url + queryParams

	return CurlInput{
		Url:        httpUrl,
		HttpMethod: httpMethod,
	}

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
