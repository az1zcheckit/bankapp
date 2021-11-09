package card

import (
	"bank/pkg/bank/types"
)

// IssueCard - Для создание/выпуска самой карты.
func IssueCard(color string, name string) types.Card {
	card := types.Card{
		ID:       1047,
		PAN:      "5058 xxxx xxxx 4047",
		Balance:  2500,
		Currency: types.TJS,
		Color:    color,
		Name:     name,
		Active:   true,
	}
	return card
}

const WithdrawLimit = 20_000_00

// Withdraw - Для снятия средств с карты.
func Withdraw(card types.Card, amount types.Money) types.Card {
	if amount <= 0 || amount > WithdrawLimit || !card.Active || card.Balance < amount {
		return card
	}
	card.Balance -= amount // снятие средств(card.Balance = card.Balance - amount)
	return card
}

const DepoLimit = 50_000_00

// Deposit - Для зачисления средств на карту.
func Deposit(card *types.Card, amount types.Money) {
	if !card.Active || amount < 0 || amount > DepoLimit {
		return
	}
	card.Balance += amount // зачисление средств(card.Balance = card.Balance + amount)
}

const BonusLimit = 5_000_00

// AddBonus - Для зачисления бонуса в размере 3% годовых на мин остаток карты.
func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {
	bonus := card.MinBalance * types.Money(percent) / 100 * types.Money(daysInMonth) / types.Money(daysInYear)
	if !card.Active || card.Balance < 0 || bonus > BonusLimit {
		return
	}
	card.Balance += bonus // зачисление бонуса(card.Balance = card.Balance + bonus)
}

// Total - Для общего расчета средств всех карт.
func Total(cards []types.Card) types.Money {
	summ := types.Money(0)
	for _, card := range cards {
		if !card.Active {
			continue // продолжение после проверки
		}
		if card.Balance <= 0 {
			continue // продолжение после проверки
		}
		summ += card.Balance
	}
	return summ
}

// Функция PaymentSources - для того чтобы выбрать источник оплаты.
func PaymentSources(cards []types.Card) []types.PaymentSource {
	var PaymentSource = []types.PaymentSource{}
	for _, card := range cards {
		if card.Active && card.Balance > 0 {
			PaymentSource = append(PaymentSource, types.PaymentSource{
				Balance: card.Balance,
				Number:  string(card.PAN),
				Type:    "card",
			})
		}
	}
	return PaymentSource
}
