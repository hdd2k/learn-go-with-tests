package main

import (
	"fmt"
	"net/http"
	"time"
)

type Racer func(string, string) (string, error)

func Race(r Racer, urlOne, urlTwo string) (string, error) {
	return r(urlOne, urlTwo)
}

func BaseRacer(urlOne, urlTwo string) (winner string, err error) {
	durationOne := measureDuration(urlOne)
	durationTwo := measureDuration(urlTwo)

	if durationOne > durationTwo {
		return urlTwo, nil
	}
	return urlOne, nil
}

func SelectRacer(urlOne, urlTwo string) (winner string, err error) {
	select {
	case <-ping(urlOne):
		return urlOne, nil
	case <-ping(urlTwo):
		return urlTwo, nil
	case <-time.After(10 * time.Second):
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
