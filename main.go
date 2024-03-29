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

	curlInput := Converter(httpConfig)

	Curl(curlInput)

}
