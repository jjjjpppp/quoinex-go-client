package models

import (
	"encoding/json"
	"sort"
)

type PriceLevels struct {
	BuyPriceLevels  [][]json.Number `json:"buy_price_levels"`
	SellPriceLevels [][]json.Number `json:"sell_price_levels"`
}

func (p *PriceLevels) GetSellPriceLevelsFloat64() [][]float64 {
	var sellFloat64 [][]float64
	for _, s := range p.SellPriceLevels {
		a, _ := s[0].Float64()
		b, _ := s[1].Float64()
		sellFloat64 = append(sellFloat64, []float64{a, b})
	}
	return sellFloat64
}

func (p *PriceLevels) GetBuyPriceLevelsFloat64() [][]float64 {
	var buyFloat64 [][]float64
	for _, buy := range p.BuyPriceLevels {
		a, _ := buy[0].Float64()
		b, _ := buy[1].Float64()
		buyFloat64 = append(buyFloat64, []float64{a, b})
	}
	return buyFloat64
}

func (p *PriceLevels) SortSellPriceLevelsByPrice(order string) [][]float64 {
	sortSell := p.GetSellPriceLevelsFloat64()
	if order == "asc" {
		sort.Slice(sortSell, func(i, j int) bool {
			return sortSell[i][0] < sortSell[j][0]
		})
	} else {
		sort.Slice(sortSell, func(i, j int) bool {
			return sortSell[i][0] > sortSell[j][0]
		})
	}
	return sortSell
}

func (p *PriceLevels) SortBuyPriceLevelsByPrice(order string) [][]float64 {
	sortBuy := p.GetBuyPriceLevelsFloat64()

	if order == "asc" {
		sort.Slice(sortBuy, func(i, j int) bool {
			return sortBuy[i][0] < sortBuy[j][0]
		})
	} else {
		sort.Slice(sortBuy, func(i, j int) bool {
			return sortBuy[i][0] > sortBuy[j][0]
		})
	}
	return sortBuy
}

func (p *PriceLevels) SortSellPriceLevelsByQuontity() [][]float64 {
	sortSell := p.GetSellPriceLevelsFloat64()
	sort.Slice(sortSell, func(i, j int) bool {
		return sortSell[i][1] > sortSell[j][1]
	})
	return sortSell
}

func (p *PriceLevels) SortBuyPriceLevelsByQuontity() [][]float64 {
	sortBuy := p.GetBuyPriceLevelsFloat64()
	sort.Slice(sortBuy, func(i, j int) bool {
		return sortBuy[i][1] > sortBuy[j][1]
	})
	return sortBuy
}
