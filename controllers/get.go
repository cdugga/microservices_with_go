package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/cdugga/microservices_with_go/data"
	"net/http"
)

func HelloWorldHandler(rw http.ResponseWriter, r *http.Request){
	c := data.GetPrices()

	var prices CryptoPrices
	json.Unmarshal(c, &prices)

	e := json.NewEncoder(rw)
	e.Encode(prices)

	fmt.Println("Return stuff")
}


type CryptoPrices struct {
	Data      []Data `json:"data"`
	Timestamp int64  `json:"timestamp"`
}
type Data struct {
	ExchangeID            string      `json:"exchangeId"`
	Rank                  string      `json:"rank"`
	BaseSymbol            string      `json:"baseSymbol"`
	BaseID                string      `json:"baseId"`
	QuoteSymbol           string      `json:"quoteSymbol"`
	QuoteID               string      `json:"quoteId"`
	PriceQuote            string      `json:"priceQuote"`
	PriceUsd              string      `json:"priceUsd"`
	VolumeUsd24Hr         string      `json:"volumeUsd24Hr"`
	PercentExchangeVolume interface{} `json:"percentExchangeVolume"`
	TradesCount24Hr       interface{} `json:"tradesCount24Hr"`
	Updated               int64       `json:"updated"`
}
