package types

// Money представляет собой денежную сумму в минимальных единицах (центы, дирамы, копейки и т.д.)
type Money int64

// Currency представляет собой код валюты.
type Currency string

// Коды валют.
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

// PAN(Payment Card Number) представляет номер карты.
type PAN string

// Card представляет собой информацию о платёжной карте.
type Card struct {
	ID         int
	PAN        PAN
	Balance    Money // использовали Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
	MinBalance Money
}

// Payment представляет собой информацию о платеже.
type Payment struct {
	ID     int
	Amount Money
}

type PaymentSource struct {
	Type    string // 'card'
	Number  string // номер вида '5058 xxxx xxxx 8888'
	Balance Money  // баланс в дирамах
}
