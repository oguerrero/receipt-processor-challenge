package common

type Item struct {
	ShortDescription  string    `json:"shortDescription"`
	Price             string   `json:"price"`
}

type Receipt struct {
	Retailer      string   `json:"retailer"` 
	PurchaseDate  string   `json:"purchaseDate"`
	PurchaseTime  string   `json:"purchaseTime"`
	Items         []Item   `json:"items"`
	Total         string  `json:"total"`
}

type ProcessedReceipt struct {
	ID string
	Points int
}
