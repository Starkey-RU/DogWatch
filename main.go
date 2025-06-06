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
	fmt.Println("shutting down...")
}

//[TODO] AlarmWatch - check if process going to shutdown by a user, not by shutdown.
//[TODO] WEB-Socket -> get a message to server
