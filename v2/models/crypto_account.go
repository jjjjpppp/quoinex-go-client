package models

type CryptoAccount struct {
	ID                       uint64  `json:"id"`
	Balance                  string  `json:"balance"`
	Address                  string  `json:"address"`
	Currency                 string  `json:"currency"`
	CurrencySymbol           string  `json:"currency_symbol"`
	PusherChannel            string  `json:"pusher_channel"`
	MinimumWithdraw          float64 `json:"minimum_withdraw"`
	LowestOfferInterestRate  string  `json:"lowest_offer_interest_rate"`
	HighestOfferInterestRate string  `json:"highest_offer_interest_rate"`
	CurrencyType             string  `json:"currency_type"`
}
