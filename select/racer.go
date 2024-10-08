package main

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondsTimeout = 10 * time.Second

// Racer compares two urls by speed of responce, timing out after 10s.
func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondsTimeout)
}

// ConfigurableRacer compares the response times of a and b, returning the fastest one.
func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
