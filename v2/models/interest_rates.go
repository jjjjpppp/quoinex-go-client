package models

type InterestRates struct {
	Bids [][]string    `json:"bids"`
	Asks []interface{} `json:"asks"`
}
