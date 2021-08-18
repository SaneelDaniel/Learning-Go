/**
	* In This program use WaitGroup, which is exactly for waiting for a bunch of go routines to finish.
	  So we are going to import sync, and then in our main, we are going to create a WaitGroup.
	  And then for every go routine that we are going to spin, so we are going to tell the WaitGroup that we added another worker,
	  and inside the go routine, we are going to signal that we are done. And lastly we are going to wait for all the go routines to be finished
**/

// Get content type of sites
package main

import (
	"fmt"
	"net/http"
	"sync"
)

func returnType(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	fmt.Printf("%s -> %s\n", url, ctype)
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			returnType(url)
			wg.Done()
		}(url)
	}
	wg.Wait()
}
