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
		t.Fatalf("Expected: %v\nActual: %v\n", expected)
	}

}
