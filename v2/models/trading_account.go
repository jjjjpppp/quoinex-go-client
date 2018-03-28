package models

type TradingAccount struct {
	ID               int    `json:"id"`
	LeverageLevel    int    `json:"leverage_level"`
	MaxLeverageLevel int    `json:"max_leverage_level"`
	Pnl              string `json:"pnl"`
	Equity           string `json:"equity"`
	Margin           string `json:"margin"`
	FreeMargin       string `json:"free_margin"`
	TraderID         int    `json:"trader_id"`
	Status           string `json:"status"`
	ProductCode      string `json:"product_code"`
	CurrencyPairCode string `json:"currency_pair_code"`
	Position         string `json:"position"`
	Balance          string `json:"balance"`
	CreatedAt        int    `json:"created_at"`
	UpdatedAt        int    `json:"updated_at"`
	PusherChannel    string `json:"pusher_channel"`
	MarginPercent    string `json:"margin_percent"`
	ProductID        int    `json:"product_id"`
	FundingCurrency  string `json:"funding_currency"`
}
