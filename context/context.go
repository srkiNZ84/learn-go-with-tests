package main

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return
		}

		fmt.Fprint(w, data)
		/* // Get the http request context
		ctx := r.Context()

		// Make a channel with buffer size of 1
		data := make(chan string, 1)

		// Fork a function to execute the "Fetch" and return the result to
		// the channel when done
		go func() {
			data <- store.Fetch()
		}()

		// Use select to have a race to see whether our work completes first
		// or we get the "cancel" command
		select {
		case d := <-data:
			fmt.Fprint(w, d)
		case <-ctx.Done():
			store.Cancel()
		}*/
	}
}
