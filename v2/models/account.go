package models

type Account struct {
	ID                       uint64 `json:"id"`
	Currency                 string `json:"currency"`
	CurrencySymbol           string `json:"currency_symbol"`
	Balance                  string `json:"balance"`
	PusherChannel            string `json:"pusher_channel"`
	LowestOfferInterestRate  string `json:"lowest_offer_interest_rate"`
	HighestOfferInterestRate string `json:"highest_offer_interest_rate"`
	ExchangeRate             string `json:"exchange_rate"`
	CurrencyType             string `json:"currency_type"`
}
