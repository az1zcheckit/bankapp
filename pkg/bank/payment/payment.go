package payment

import (
	"bank/pkg/bank/types"
)

// Max возвращает максимальный платёж.
func Max(payments []types.Payment) types.Payment {
	maxPayment := payments[0]
	for _, payment := range payments {
		if maxPayment.Amount < payment.Amount {
			maxPayment = payment
		}
	}
	return maxPayment
}