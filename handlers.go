package main

import (
	"fmt"
	"net/http"
)

func GetCode(w http.ResponseWriter, req *http.Request) {
	i := <-randIntChanel

	mu.RLock()
	code := queueOrder[i]
	mu.RUnlock()

	mu.Lock()
	statusCode := topOrder[code]
	statusCode.rate = statusCode.rate + 1
	topOrder[code] = statusCode
	mu.Unlock()

	fmt.Fprintf(w, "%s", code)
}

func GetRateCodes(w http.ResponseWriter, req *http.Request) {
	mu.RLock()
	for k, v := range topOrder {
		if v.rate > 0 {
			fmt.Fprintf(w, "%s - %d\n", k, v.rate)
		}
	}
	mu.RUnlock()
}
