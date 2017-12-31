package request

import (
	"github.com/gugadev/cryptogo/models"
	"github.com/levigross/grequests"
)

// URI Base url of the MarketCap API
const URI = "https://api.coinmarketcap.com/v1/ticker/"

/*
GetInfo Get info of a crypto currency and
store the information inside a struct
*/
func GetInfo(coin string) []models.Coin {
	var result []models.Coin
	options := &grequests.RequestOptions{
		Params: map[string]string{
			"convert": "EUR",
		},
	}
	response, err := grequests.Get(URI+coin, options)
	if err != nil {
		panic(err)
	}
	response.JSON(&result)
	return result
}
