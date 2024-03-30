package main

import (
	"fmt"
	"os/exec"
)

func Curl(input CurlInput, cliFlags CommandLineFlags) {

	httpMethod := input.HttpMethod
	url := input.Url

	if cliFlags.DebugMode {
		fmt.Printf("http method: %v \n", httpMethod)
		fmt.Printf("url: %v \n", url)
	}

	cmd := exec.Command("curl", "-X", httpMethod, url)

	out, err := cmd.Output()

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(string(out))

}
