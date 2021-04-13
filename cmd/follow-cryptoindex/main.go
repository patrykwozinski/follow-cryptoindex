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

	gecko := coingecko.NewClient(httpClient)

	coins, err := gecko.CoinsMarket("USD", make([]string, 0), "", 1, 1, true, make([]string, 0))

	if err != nil {
		log.Fatal("error xd")
	}

	for _, s := range *coins {
		fmt.Println(s.Name)
		fmt.Println(s.MarketCap)

		for _, price := range s.SparklineIn7d.Price {
			fmt.Println(price)
		}
	}
}
