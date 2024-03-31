package internal

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

var availableMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}

func ValidateInput(httpConfig HttpRequestConfig) []error {

	var errorsList []error = make([]error, 0)

	if httpConfig.Url == "" {
		errorsList = append(errorsList, errors.New("Missing URL in Input"))
	}

	method := httpConfig.Method

	if !isValidHttpMethod(method) {
		wrongMethodMessage := fmt.Sprintf("Invalid input, got %v, expected one of %v \n", method, availableMethods)

		errorsList = append(errorsList, errors.New(wrongMethodMessage))
	}

	//TODO: add validation for path params

	return errorsList
}

func isValidHttpMethod(method string) bool {
	return slices.Contains(availableMethods, strings.ToUpper(method))
}
