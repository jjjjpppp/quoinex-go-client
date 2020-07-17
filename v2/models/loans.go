package models

type Loans struct {
	Models      []*Loan `json:"models"`
	CurrentPage int     `json:"current_page"`
	TotalPages  int     `json:"total_pages"`
}

type Loan struct {
	ID           uint64 `json:"id"`
	Quantity     string `json:"quantity"`
	Rate         string `json:"rate"`
	CreatedAt    uint   `json:"created_at"`
	LenderID     uint   `json:"lender_id"`
	BorrowerID   uint   `json:"borrower_id"`
	Status       string `json:"status"`
	Currency     string `json:"currency"`
	FundReloaned bool   `json:"fund_reloaned"`
}
