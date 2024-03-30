package main

func ProcessFlags(flags []string) CommandLineFlags {

	var flagsMap = make(map[string]bool)

	for _, flag := range flags {
		flagsMap[flag] = true
	}

	_, isDebug := flagsMap["--debug"]

	return CommandLineFlags{
		DebugMode: isDebug,
	}
}
