package data

import (
	"io/ioutil"
	"net/http"
)

func GetPrices() []byte {

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, "https://api.coincap.io/v2/markets?exchangeId=poloniex&limit=15", nil)

	resp, err := client.Do(req)

	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return nil
	}

	return body


//	return &CryptoPrices{
//		Data:      []Data{{
//			ExchangeID:            "poloniex",
//			Rank:                  "2",
//			BaseSymbol:            "BTC",
//			BaseID:                "bitcoin",
//			QuoteSymbol:           "USDT",
//			QuoteID:               "tether",
//			PriceQuote:            "9662.8669251000000000",
//			PriceUsd:              "9675.7322559662744161",
//			VolumeUsd24Hr:         "8306299.8195151428923058",
//			PercentExchangeVolume: "23.8263238203500179",
//			TradesCount24Hr:       "19720",
//			Updated:               1595704652095,
//		}},
//		Timestamp: 1595704652095,
//	}
}



