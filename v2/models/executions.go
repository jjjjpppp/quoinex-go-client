package models

type Executions struct {
	Models []struct {
		ID        int    `json:"id"`
		Quantity  string `json:"quantity"`
		Price     string `json:"price"`
		TakerSide string `json:"taker_side"`
		CreatedAt int    `json:"created_at"`
	} `json:"models"`
	CurrentPage int `json:"current_page"`
	TotalPages  int `json:"total_pages"`
}
