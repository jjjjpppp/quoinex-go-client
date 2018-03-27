package models

type Loans struct {
	Models      []*Loan `json:"models"`
	CurrentPage int     `json:"current_page"`
	TotalPages  int     `json:"total_pages"`
}

type Loan struct {
	ID           int    `json:"id"`
	Quantity     string `json:"quantity"`
	Rate         string `json:"rate"`
	CreatedAt    int    `json:"created_at"`
	LenderID     int    `json:"lender_id"`
	BorrowerID   int    `json:"borrower_id"`
	Status       string `json:"status"`
	Currency     string `json:"currency"`
	FundReloaned bool   `json:"fund_reloaned"`
}
