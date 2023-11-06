package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jpillora/overseer"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	// Create an overseer instance
	overseer.Run(overseer.Config{
		Program: func(state overseer.State) {
			// This code will be executed after an update
			http.Serve(state.Listener, nil)
		},
	})

	// The overseer.Run() function will block the main program and handle updates
	os.Exit(0)
}
