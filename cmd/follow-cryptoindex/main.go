package main

import (
	"fmt"
	coingecko "github.com/superoo7/go-gecko/v3"
	"log"
	"net/http"
	"time"
)

func main() {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	coins, err := coingecko.NewClient(httpClient).CoinsMarket("USD", make([]string, 0), "", 3, 1, true, make([]string, 0))

	if err != nil {
		log.Fatal("error xd")
	}

	for _, s := range *coins {
		fmt.Println(s.Name)
		fmt.Println(s.MarketCap)
	}
}
