package main

import (
	"log"
	"net/http"
	"sync"
	"time"
)

type status struct {
	state bool
	rate  int
}

const SizeStr = 2
const sizeQueue = 50

var queueOrder [sizeQueue][SizeStr]byte
var mu sync.RWMutex

var topOrder = make(map[[SizeStr]byte]status)

var randIntChanel = make(chan int)
var randStrChanel = make(chan [SizeStr]byte)

func main() {

	go IntGenerator(randIntChanel)

	go StrGenerator(randStrChanel)

	initQueue()

	go changer()

	http.HandleFunc("/request", GetCode)
	http.HandleFunc("/admin/request", GetRateCodes)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func initQueue() {
	for k, _ := range queueOrder {
		for r := true; r; {
			code := <-randStrChanel
			if _, ok := topOrder[code]; !ok {
				topOrder[code] = status{true, 0}
				queueOrder[k] = code
				r = false
			}
		}
	}
}

func changer() {
	timer := time.Tick(200 * time.Millisecond)
	for {
		<-timer
		i := <-randIntChanel

		for r := true; r; {
			code := <-randStrChanel

			mu.RLock()
			statusCode, ok := topOrder[code]
			mu.RUnlock()

			if !ok || !statusCode.state {

				statusCode.state = true

				mu.Lock()
				topOrder[code] = statusCode

				oldCode := queueOrder[i]
				topOrder[oldCode] = status{false, topOrder[oldCode].rate}
				queueOrder[i] = code
				mu.Unlock()

				r = false
			}
		}

	}
}
