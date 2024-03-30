package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {

	inputAsBytes, err := io.ReadAll(os.Stdin)

	if err != nil {
		panic(err.Error())
	}

	httpConfig := HttpRequestConfig{}

	if err := json.Unmarshal(inputAsBytes, &httpConfig); err != nil {
		fmt.Println("error parsing json")
		panic(err.Error())
	}

	errors := ValidateInput(httpConfig)

	if len(errors) > 0 {

		for _, error := range errors {
			fmt.Println(error.Error())
		}

		panic("Invalid Inputs")
	}

	curlInput := Converter(httpConfig)

	Curl(curlInput)

}
