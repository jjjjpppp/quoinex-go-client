package models

type Product struct {
	ID                  uint64 `json:"id"`
	ProductType         string `json:"product_type"`
	Code                string `json:"code"`
	Name                string `json:"name"`
	MarketAsk           string `json:"market_ask"`
	MarketBid           string `json:"market_bid"`
	Indicator           int    `json:"indicator"`
	Currency            string `json:"currency"`
	CurrencyPairCode    string `json:"currency_pair_code"`
	Symbol              string `json:"symbol"`
	FiatMinimumWithdraw string `json:"fiat_minimum_withdraw"`
	PusherChannel       string `json:"pusher_channel"`
	TakerFee            string `json:"taker_fee"`
	MakerFee            string `json:"maker_fee"`
	LowMarketBid        string `json:"low_market_bid"`
	HighMarketAsk       string `json:"high_market_ask"`
	Volume24H           string `json:"volume_24h"`
	LastPrice24H        string `json:"last_price_24h"`
	LastTradedPrice     string `json:"last_traded_price"`
	LastTradedQuantity  string `json:"last_traded_quantity"`
	QuotedCurrency      string `json:"quoted_currency"`
	BaseCurrency        string `json:"base_currency"`
	ExchangeRate        string `json:"exchange_rate"`
}
