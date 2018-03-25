package models

type Trade struct {
	ID               int    `json:"id"`
	CurrencyPairCode string `json:"currency_pair_code"`
	Status           string `json:"status"`
	Side             string `json:"side"`
	MarginUsed       string `json:"margin_used"`
	OpenQuantity     string `json:"open_quantity"`
	CloseQuantity    string `json:"close_quantity"`
	Quantity         string `json:"quantity"`
	LeverageLevel    int    `json:"leverage_level"`
	ProductCode      string `json:"product_code"`
	ProductID        int    `json:"product_id"`
	OpenPrice        string `json:"open_price"`
	ClosePrice       string `json:"close_price"`
	TraderID         int    `json:"trader_id"`
	OpenPnl          string `json:"open_pnl"`
	ClosePnl         string `json:"close_pnl"`
	Pnl              string `json:"pnl"`
	StopLoss         string `json:"stop_loss"`
	TakeProfit       string `json:"take_profit"`
	FundingCurrency  string `json:"funding_currency"`
	CreatedAt        int    `json:"created_at"`
	UpdatedAt        int    `json:"updated_at"`
	CloseFee         string `json:"close_fee"`
	TotalInterest    string `json:"total_interest"`
	DailyInterest    string `json:"daily_interest"`
}
