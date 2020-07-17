package models

type Orders struct {
	Models      []*Order `json:"models"`
	CurrentPage int      `json:"current_page"`
	TotalPages  int      `json:"total_pages"`
}

type Order struct {
	ID                   uint64      `json:"id"`
	OrderType            string      `json:"order_type"`
	Quantity             string      `json:"quantity"`
	DiscQuantity         string      `json:"disc_quantity"`
	IcebergTotalQuantity string      `json:"iceberg_total_quantity"`
	Side                 string      `json:"side"`
	FilledQuantity       string      `json:"filled_quantity"`
	Price                float64     `json:"price"`
	CreatedAt            uint        `json:"created_at"`
	UpdatedAt            uint        `json:"updated_at"`
	Status               string      `json:"status"`
	LeverageLevel        uint        `json:"leverage_level"`
	SourceExchange       string      `json:"source_exchange"`
	ProductID            uint        `json:"product_id"`
	MarginType           interface{} `json:"margin_type"`
	TakeProfit           interface{} `json:"take_profit"`
	StopLoss             interface{} `json:"stop_loss"`
	TradingType          string      `json:"trading_type"`
	ProductCode          string      `json:"product_code"`
	FundingCurrency      string      `json:"funding_currency"`
	CryptoAccountID      interface{} `json:"crypto_account_id"`
	CurrencyPairCode     string      `json:"currency_pair_code"`
	AveragePrice         float64     `json:"average_price"`
	Target               string      `json:"target"`
	OrderFee             interface{} `json:"order_fee"`
	SourceAction         string      `json:"source_action"`
	UnwoundTradeID       interface{} `json:"unwound_trade_id"`
	TradeID              interface{} `json:"trade_id"`
	ClientOrderID        interface{} `json:"client_order_id"`
}
