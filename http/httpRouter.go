package router

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type Router struct {
	config *RouterConfig
}

type RouterConfig struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Target1Url   string
	Target2Url   string
}

func NewRouter(config *RouterConfig) *Router {
	return &Router{
		config: config,
	}
}

func (r *Router) Start() error {
	http.HandleFunc("/", r.handler)
	return http.ListenAndServe(r.config.Addr, nil)
}

func (router *Router) handler(w http.ResponseWriter, r *http.Request) {
	resultCh := make(chan string, 2)
	go sendRequest(r, router.config.Target1Url, resultCh)
	go sendRequest(r, router.config.Target2Url, resultCh)

	result1 := <-resultCh
	result2 := <-resultCh
	close(resultCh)

	// fmt.Printf("Result 1: %s\n", result1)
	// fmt.Printf("Result 2: %s\n", result2)
	fmt.Printf("Are responses equal: %t\n", result1 == result2)
}

func sendRequest(r *http.Request, url string, resultCh chan<- string) {
	// Create a new HTTP client
	client := &http.Client{}

	// Create a new request based on the original request
	req, err := http.NewRequest(r.Method, url, r.Body)
	if err != nil {
		// Handle error
		resultCh <- err.Error()
		return
	}
	req.Header = r.Header.Clone()

	resp, err := client.Do(req)
	if err != nil {
		resultCh <- err.Error()
		return
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		// Handle error
		resultCh <- err.Error()
		return
	}
	// Send the result through the channel
	resultCh <- string(responseBody)
}
