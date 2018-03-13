package models

type PriceLevels struct {
	BuyPriceLevels  [][]string `json:"buy_price_levels"`
	SellPriceLevels [][]string `json:"sell_price_levels"`
}
