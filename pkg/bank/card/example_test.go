package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

// Тест снятия денег с карты.
func ExampleWithdraw_positive() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: true}, 10_000_00)
	fmt.Println(result.Balance)
	// Output: 1000000
}
func ExampleWithdraw_noMoney() {
	result := Withdraw(types.Card{Balance: 5_000_00, Active: true}, 10_000_00)
	fmt.Println(result.Balance)
	// Output: 500000
}
func ExampleWithdraw_inActive() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: false}, 10_000_00)
	fmt.Println(result.Active)
	// Output: false
}
func ExampleWithdraw_drawLimit() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: true}, 30_000_00)
	fmt.Println(result.Balance)
	// Output: 2000000
}

// Тест зачисления денег на карту.
func ExampleDeposit_positive() {
	card := types.Card{Balance: 10_000_00, Active: true}
	Deposit(&card, 5_000_00)
	fmt.Println(card.Balance)
	// Output: 1500000
}

func ExampleDeposit_inActive() {
	card := types.Card{Balance: 10_000_00, Active: false}
	Deposit(&card, 20_000_00)
	fmt.Println(card.Active)
	// Output: false
}
func ExampleDeposit_depoLimit() {
	card := types.Card{Balance: 40_000_00, Active: true}
	Deposit(&card, 40_000_00)
	fmt.Println(card.Balance)
	// Output: 8000000
}

// Тест зачисления бонуса на карту.
func ExampleAddBonus_positive() {
	card := types.Card{Balance: 10_000_00, MinBalance: 10_000_00, Active: true}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: 1002465
}
func ExampleAddBonus_inActive() {
	card := types.Card{Balance: 10_000_00, MinBalance: 10_000_00, Active: false}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Active)
	// Output: false
}
func ExampleAddBonus_negative() {
	card := types.Card{Balance: -10_000_00, MinBalance: 10_000_00, Active: true}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: -1000000
}
func ExampleAddBonus_bonusLimit() {
	card := types.Card{Balance: 2_00_000_00, MinBalance: 2_00_000_00, Active: true}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: 20049315
}

// Тест для общего расчета всех карт.
func ExampleTotal() {
	cards := []types.Card{
		{
			Balance: 10,
			Active:  true,
		},
		{
			Balance: 10,
			Active:  true,
		},
	}
	fmt.Println(Total(cards))
	// Output: 20
}

// Тест для выбора источников оплаты из карт.
func ExamplePaymentSources() {
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
	sources := PaymentSources(cards)
	for _, source := range sources {
		fmt.Println(source.Number)
	}

	// Output :
	// 5058 2702 8808 0001
	// 5058 2702 8808 0004
}
