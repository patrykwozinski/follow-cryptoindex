package main

import (
	"fmt"
	coingecko "github.com/superoo7/go-gecko/v3"
	"net/http"
	"time"
)

func main() {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	coingecko.NewClient(httpClient)

	fmt.Printf("Buy high, sell low :)")
}
