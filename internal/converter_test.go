package internal

import (
	"testing"
)

func TestSubstitutePathsParams(t *testing.T) {

	pathParams := make(map[string]string)
	pathParams["city"] = "london"

	urlWithPlaceholders := "https://holiday-website.com/:city"

	actual := substitutePathParams(urlWithPlaceholders, pathParams)

	expected := "https://holiday-website.com/london"
	if actual != expected {
		t.Fatalf("Expected: %v\nActual: %v\n", expected, actual)
	}

}

func TestGenerateQueryParams_ParamsProvided(t *testing.T) {
	queryParams := make(map[string][]string)
	queryParams["name"] = []string{"foo", "bar"}

	actual := generateQueryParamsString(queryParams)

	expected := "?name=foo&name=bar"

	if actual != expected {
		t.Fatalf("Expected: %v\nActual: %v\n", expected, actual)
	}

}

func TestGenerateQueryParams_EmptyParams(t *testing.T) {
	queryParams := make(map[string][]string)

	actual := generateQueryParamsString(queryParams)

	expected := ""

	if actual != expected {
		t.Fatalf("Expected: %v\nActual: %v\n", expected, actual)
	}

}

func TestGenerateBasicAuthHeaderValue(t *testing.T) {
	credentials := BasicAuth{
		Username: "foo",
		Password: "bar",
	}

	actual := generateBasicAuthHeader(credentials)

	expected := "Basic Zm9vOmJhcg=="

	if actual != expected {
		t.Fatalf("Expected: %v\nActual: %v\n", expected, actual)
	}

}
