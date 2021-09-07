package main

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecTimeout = 10 * time.Second

// type Racer func(string, string) (string, error)

func Racer(urlOne, urlTwo string) (string, error) {
	return ConfigurableRacer(urlOne, urlTwo, tenSecTimeout)
}

func BaseRacer(urlOne, urlTwo string) (winner string, err error) {
	durationOne := measureDuration(urlOne)
	durationTwo := measureDuration(urlTwo)

	if durationOne > durationTwo {
		return urlTwo, nil
	}
	return urlOne, nil
}

func ConfigurableRacer(urlOne, urlTwo string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(urlOne):
		return urlOne, nil
	case <-ping(urlTwo):
		return urlTwo, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s ad %s", urlOne, urlTwo)
	}
}

func measureDuration(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
