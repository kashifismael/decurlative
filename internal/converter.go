package internal

import (
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

	queryParams := generateQueryParamsString(httpConfig.QueryParams)

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
