package utils

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func ListenForKeyboardInterrupted(fn func()) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		if fn != nil {
			fn()
		}
		fmt.Println("\nShutdown the server, bye...")
		os.Exit(0)
	}()
}
