package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("Starting reporter...")

	setupCloseHandler()

	c := NewConsumer()

	StartServer(c)
}

func setupCloseHandler() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Goodbye")
		os.Exit(0)
	}()
}
