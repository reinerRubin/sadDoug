package main

import (
	"os"
	"os/signal"
	"syscall"

	saddoug "github.com/sadDoug"
)

func main() {
	appEnv, err := saddoug.NewAppEnv()
	if err != nil {
		panic(err)
	}

	messageSource, err := saddoug.NewBNWThread("H17XM7")
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

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
}
