package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

const (
	concurrency     = 100
	continuationSec = 10
)

func main() {

	var wg sync.WaitGroup
	wg.Add(concurrency * continuationSec)

	for i := 0; i < continuationSec; i++ {
		fmt.Println(i)

		for j := 0; j < concurrency; j++ {
			go func(i int) {
				defer wg.Done()

				resp, err := http.Get("http://localhost:8000")
				if err != nil {
					log.Fatalln(err)
				}

				_, err = io.ReadAll(resp.Body)
				if err != nil {
					log.Fatalln(err)
				}
				defer resp.Body.Close()
			}(j)
		}

		time.Sleep(time.Second * 1)
	}

	wg.Wait()
}
