package models

type TradingAccount struct {
	ID               uint64 `json:"id"`
	LeverageLevel    uint   `json:"leverage_level"`
	MaxLeverageLevel uint   `json:"max_leverage_level"`
	Pnl              string `json:"pnl"`
	Equity           string `json:"equity"`
	Margin           string `json:"margin"`
	FreeMargin       string `json:"free_margin"`
	TraderID         uint   `json:"trader_id"`
	Status           string `json:"status"`
	ProductCode      string `json:"product_code"`
	CurrencyPairCode string `json:"currency_pair_code"`
	Position         string `json:"position"`
	Balance          string `json:"balance"`
	CreatedAt        uint   `json:"created_at"`
	UpdatedAt        uint   `json:"updated_at"`
	PusherChannel    string `json:"pusher_channel"`
	MarginPercent    string `json:"margin_percent"`
	ProductID        uint   `json:"product_id"`
	FundingCurrency  string `json:"funding_currency"`
}
