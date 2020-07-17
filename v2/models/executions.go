package models

type Executions struct {
	Models      []*ExecutionsModels `json:"models"`
	CurrentPage int                 `json:"current_page"`
	TotalPages  int                 `json:"total_pages"`
}

type ExecutionsModels struct {
	ID        uint64 `json:"id"`
	Quantity  string `json:"quantity"`
	Price     string `json:"price"`
	TakerSide string `json:"taker_side"`
	MySide    string `json:"my_side"`
	CreatedAt uint   `json:"created_at"`
}
