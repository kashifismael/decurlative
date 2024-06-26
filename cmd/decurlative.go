package main

import (
	"decurlative/internal"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {

	cliArgs := os.Args[1:]

	cliFlags := internal.ProcessFlags(cliArgs)

	inputAsBytes, err := io.ReadAll(os.Stdin)

	if err != nil {
		panic(err.Error())
	}

	httpConfig := internal.HttpRequestConfig{}

	if err := json.Unmarshal(inputAsBytes, &httpConfig); err != nil {
		fmt.Println("error parsing json")
		panic(err.Error())
	}

	errors := internal.ValidateInput(httpConfig)

	if len(errors) > 0 {

		for _, error := range errors {
			fmt.Println(error.Error())
		}

		panic("Invalid Inputs")
	}

	curlInput := internal.Converter(httpConfig)

	internal.Curl(curlInput, cliFlags)

}
