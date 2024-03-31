package internal

import (
	"fmt"
	"os/exec"
)

func Curl(input CurlInput, cliFlags CommandLineFlags) {

	httpMethod := input.HttpMethod
	url := input.Url

	curlArgs := []string{"-X", httpMethod, url}

	for k, v := range input.Headers {
		curlArgs = append(curlArgs, "-H")
		curlArgs = append(curlArgs, fmt.Sprintf("%v: %v", k, v))
	}

	if cliFlags.DebugMode {
		fmt.Printf("curl args: %v \n", curlArgs)
	}

	cmd := exec.Command("curl", curlArgs...)

	out, err := cmd.Output()

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(string(out))

}
