package main

import (
	"bank/pkg/bank/card"
	"bank/pkg/bank/types"
	"fmt"
)

func main() {
	cards := []types.Card{
		{
			Balance: 1000,
			Active:  true,
			PAN:     "5058 2702 8808 0001",
		},
		{
			Balance: 1000,
			Active:  false,
			PAN:     "5058 2702 8808 0002",
		},
		{
			Balance: -1000,
			Active:  true,
			PAN:     "5058 2702 8808 0003",
		},
		{
			Balance: 2000,
			Active:  true,
			PAN:     "5058 2702 8808 0004",
		},
	}
	sources := card.PaymentSources(cards)
	for _, source := range sources {
		fmt.Println((source.Number))
	}
}
