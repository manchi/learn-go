package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("webrequests")

	myUrl := "https://www.msn.com/en-in"

	parsedUrl, err := url.Parse(myUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println("parsedUrl", parsedUrl)
	fmt.Println("Scheme", parsedUrl.Scheme)
	fmt.Println("Host:", parsedUrl.Host)
	fmt.Println("Port", parsedUrl.Port())
	fmt.Println("Path", parsedUrl.Path)

	anotherUrl := &url.URL{
		Scheme: "https",
		Host:   "msn.com",
		Path:   "en-in",
	}
	fmt.Println("another url:", anotherUrl)

	res, err := http.Get(parsedUrl.String())
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	fmt.Printf("the response type is :%T\n", res)
	databytes, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("StatusCode:", res.StatusCode)
	fmt.Println("ContentLength:", res.ContentLength)

	// parse bytes - method 1
	content := string(databytes)
	fmt.Println("content:", len(content))

	// parse bytes - method 2
	var respString strings.Builder
	byteCount, _ := respString.Write(databytes)
	fmt.Println("byteCount:", byteCount)
	fmt.Println("respString len", respString.Len())

}
