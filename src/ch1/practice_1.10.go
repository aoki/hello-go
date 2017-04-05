package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	header := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch, header)
	}
	for range os.Args[1:] {
		fmt.Printf("%s\t%s\n", <-ch, <-header)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string, header chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	defer resp.Body.Close()

	fileName := regexp.MustCompile(`^http(s)?://`).ReplaceAllString(url, "") + ".res"
	fileName = regexp.MustCompile(`/`).ReplaceAllString(fileName, "-")
	fmt.Println(fileName)
	dest, err := os.Create(fileName)
	if err != nil {
		ch <- fmt.Sprint(err)
	}
	defer dest.Close()

	nbytes, err := io.Copy(dest, resp.Body)

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
	header <- fmt.Sprintf("%s\t%s", resp.Header["Expires"], resp.Header["Cache-Control"])

}
