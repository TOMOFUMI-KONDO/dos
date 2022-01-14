package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

var (
	url         string
	concurrency int
)

func init() {
	flag.IntVar(&concurrency, "c", 100, "concurrency")
	flag.Parse()
	fmt.Println(concurrency)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("specify server url")
	}
	url = os.Args[1]

	var wg sync.WaitGroup
	wg.Add(concurrency)

	fmt.Println("start")
	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()

			resp, err := http.Get(url)
			if err != nil {
				log.Println(err)
			}

			_, err = io.ReadAll(resp.Body)
			if err != nil {
				log.Println(err)
			}
			defer resp.Body.Close()
		}()
	}

	wg.Wait()
	fmt.Println("done")
}
