package main

import (
	"fmt"
	"os/exec"
)

func Curl(input CurlInput) {

	httpMethod := input.HttpMethod
	url := input.Url

	cmd := exec.Command("curl", "-X", httpMethod, url)

	out, err := cmd.Output()

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(string(out))

}
