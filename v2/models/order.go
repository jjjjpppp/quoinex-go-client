package models

type Order struct {
	ID                   int             `json:"id"`
	OrderType            string          `json:"order_type"`
	Quantity             string          `json:"quantity"`
	DiscQuantity         string          `json:"disc_quantity"`
	IcebergTotalQuantity string          `json:"iceberg_total_quantity"`
	Side                 string          `json:"side"`
	FilledQuantity       string          `json:"filled_quantity"`
	Price                string          `json:"price"`
	CreatedAt            int             `json:"created_at"`
	UpdatedAt            int             `json:"updated_at"`
	Status               string          `json:"status"`
	LeverageLevel        int             `json:"leverage_level"`
	SourceExchange       string          `json:"source_exchange"`
	ProductID            int             `json:"product_id"`
	ProductCode          string          `json:"product_code"`
	FundingCurrency      string          `json:"funding_currency"`
	CurrencyPairCode     string          `json:"currency_pair_code"`
	OrderFee             string          `json:"order_fee"`
	Executions           OrderExecutions `json:"executions"`
}

type OrderExecutions []struct {
	ID        int    `json:"id"`
	Quantity  string `json:"quantity"`
	Price     string `json:"price"`
	TakerSide string `json:"taker_side"`
	MySide    string `json:"my_side"`
	CreatedAt int    `json:"created_at"`
}
