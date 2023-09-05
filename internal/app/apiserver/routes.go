package apiserver

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

var (
	id int
)

func (s *APIServer) handleHello() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// Create a wait group to synchronize the response
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			id++
			text := fmt.Sprintf("hello %v\n", id)
			s.logger.Info(text)
			io.WriteString(w, "Hello")
		}()
		// Wait for the asynchronous task to complete
		wg.Wait()
	}
}

func (s *APIServer) handleInfo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Create a wait group to synchronize the response
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			id++
			text := fmt.Sprintf("info %v\n", id)
			s.logger.Info(text)
			io.WriteString(w, "Info")
		}()
		// Wait for the asynchronous task to complete
		wg.Wait()
	}
}
