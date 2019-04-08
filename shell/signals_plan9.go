// +build plan9

package shell

import (
	"os"
	"os/signal"
	"syscall"
)

// SignalHandler is an internal function to capture and handle OS signals (eg SIGTERM).
func SignalHandler(interactive bool) {
	c := make(chan os.Signal, 1)

	if Interactive {
		// Interactive, so we will handle stop
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	} else {
		// Non-interactive, so lets ignore the stop signal and let the OS / calling shell manage that for us
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	}

	go func() {
		for {
			sig := <-c
			switch sig.String() {

			case syscall.SIGINT.String():
				sigint(interactive)

			case syscall.SIGTERM.String():
				sigterm(interactive)

			default:
				os.Stderr.WriteString("Unhandled signal: " + sig.String())
			}
		}
	}()
}
