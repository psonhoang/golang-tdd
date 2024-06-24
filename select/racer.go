package main

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	// select
	// -> helps you wait on multiple channels
	// -> sometimes you might want to include time.After to prevent your system blocking forevers
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

// func measureResponseTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)
// 	return time.Since(start)
// }

func ping(url string) chan struct{} {
	ch := make(chan struct{})	// always make channels because if we define var ch channel struct{} it will be defaulted to nil
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}