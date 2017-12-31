package models

/*
Coin represent a response of
the marketcap request
*/
type Coin struct {
	Name     string
	Symbol   string
	Rank     string
	PriceUSD string `json:"price_usd"`
	PriceEUR string `json:"price_eur"`
}
