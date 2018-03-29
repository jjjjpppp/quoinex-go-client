package models

type Trades struct {
	Models      []*Trade `json:"models"`
	CurrentPage int      `json:"current_page"`
	TotalPages  int      `json:"total_pages"`
}
