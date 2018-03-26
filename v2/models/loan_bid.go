package models

type LoanBid struct {
	ID             int    `json:"id"`
	BidaskType     string `json:"bidask_type"`
	Quantity       string `json:"quantity"`
	Currency       string `json:"currency"`
	Side           string `json:"side"`
	FilledQuantity string `json:"filled_quantity"`
	Status         string `json:"status"`
	Rate           string `json:"rate"`
	UserID         int    `json:"user_id"`
}
