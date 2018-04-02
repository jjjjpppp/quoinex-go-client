package models

import (
	"encoding/json"
)

type PriceLevels struct {
	BuyPriceLevels       [][]json.Number `json:"buy_price_levels"`
	SellPriceLevels      [][]json.Number `json:"sell_price_levels"`
	BuyPriceLevelsFloat  [][]float64
	SellPriceLevelsFloat [][]float64
}

func (p *PriceLevels) ConvertToFloat64() {
	for _, pLevel := range p.BuyPriceLevels {
		a, _ := pLevel[0].Float64()
		b, _ := pLevel[1].Float64()
		p.BuyPriceLevelsFloat = append(p.BuyPriceLevelsFloat, []float64{a, b})
	}
	for _, pLevel := range p.SellPriceLevels {
		a, _ := pLevel[0].Float64()
		b, _ := pLevel[1].Float64()
		p.SellPriceLevelsFloat = append(p.SellPriceLevelsFloat, []float64{a, b})
	}
}
