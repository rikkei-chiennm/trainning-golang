package main

import (
	"fmt"
	"net/http"
	"sync"
)

// dung waitGroup chuong trinh doi trong gourotine
var wg sync.WaitGroup

func main() {
	listSite := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://stackoverflow.com",
	}
	wg.Add(len(listSite))
	for _, link := range listSite {
		go checkSite(link)
	}
	wg.Wait()

	// timesleep la lam chuong trinh ngu trong mot thoi gian nhat dinh
	//time.Sleep(10 * time.Second)

}
func checkSite(link string) {
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println(link + "is down!")
		wg.Done()
		return
	}
	fmt.Println(link + "is OK!" + resp.Status)
	wg.Done()
}
