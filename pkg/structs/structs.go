package structs

type Model struct {
	Price map[string][]string `json:"price"`
}

type Result struct {
	QuoteType string `json:"quoteType"`
}
