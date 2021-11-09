package payment

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleMax() {
	payments := []types.Payment{
		{
			ID:     1,
			Amount: 30,
		},
		{
			ID:     2,
			Amount: 30,
		},
	}

	max := Max(payments)
	fmt.Println(max.ID)

	//Output: 1
}
