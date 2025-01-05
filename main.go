package main

import (
	"fmt"

	"github.com/Yokomi422/go-scraper/source"
)

func main() {
	url := "https://pkg.go.dev/net/http"

	body, err := source.Body(url)
	if err != nil {
		panic(err)
	}

	fmt.Println(body)
}
