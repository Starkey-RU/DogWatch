package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	StartCronTask()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	fmt.Println("Shutting down...")
}
