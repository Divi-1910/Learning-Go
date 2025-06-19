package main

import (
	"fmt"
	"net/url"
)

func main() {
	// format of url :-  [scheme:://[userinfo@]host[:port][/path][?query][#fragment]

	rawUrl := "https://example.com:8080/path?query=param#fragment"

	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		fmt.Println("Parse URL error:", err)
		return
	}
	fmt.Println("Scheme : ", parsedUrl.Scheme)
	fmt.Println("Host : ", parsedUrl.Host)
	fmt.Println("Path : ", parsedUrl.Path)
	fmt.Println("Query : ", parsedUrl.Query())
	fmt.Println("Port :", parsedUrl.Port())
	fmt.Println("Fragment : ", parsedUrl.Fragment)
	fmt.Println("parsed Url : ", parsedUrl)

	rawUrl = "https://example.com:8080/path?name=John&age=30"
	parsedUrl, err = url.Parse(rawUrl)
	if err != nil {
		fmt.Println("Parse URL error:", err)
		return
	}

	queryParams := parsedUrl.Query()
	fmt.Println("Query : ", queryParams) // this is a map
	fmt.Println("Name : ", queryParams["name"])
	fmt.Println("Age : ", queryParams["age"])

	// Building URL
	baseUrl := &url.URL{
		Scheme: "https",
		Host:   "example.com",
		Path:   "/path",
	}

	query := baseUrl.Query()
	query.Set("name", "John")
	baseUrl.RawQuery = query.Encode()
	fmt.Println("Query : ", baseUrl)

	fmt.Println("Base : ", baseUrl.String())

	values := url.Values{}

	values.Add("name", "John")
	values.Add("age", "30")
	baseUrl.RawQuery = values.Encode()

	fmt.Println("Query : ", baseUrl)

}
