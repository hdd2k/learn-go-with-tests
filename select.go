package main

import (
	"net/http"
	"time"
)

type Racer func(string, string) string

func Race(r Racer, urlOne, urlTwo string) string {
	return r(urlOne, urlTwo)
}

func BaseRacer(urlOne, urlTwo string) (winner string) {
	durationOne := measureDuration(urlOne)
	durationTwo := measureDuration(urlTwo)

	if durationOne > durationTwo {
		return urlTwo
	}
	return urlOne
}

func measureDuration(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
