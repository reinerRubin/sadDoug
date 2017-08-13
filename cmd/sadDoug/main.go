package main

import (
	saddoug "github.com/sadDoug"
)

func main() {
	appEnv, err := saddoug.NewAppEnv()
	if err != nil {
		panic(err)
	}

	messageSource, err := saddoug.NewBNWThread("EA3W6L")
	if err != nil {
		panic(err)
	}

	messageSaver, err := saddoug.NewMessageSaver(appEnv, messageSource)
	if err != nil {
		panic(err)
	}

	if err := messageSaver.Run(); err != nil {
		panic(err)
	}
}
