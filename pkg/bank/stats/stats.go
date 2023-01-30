package stats

import ""


// AVG рассчитывает среднюю сумму платежа.
func AVG(payments []types.Payment) types.Money {

	summPays := types.Money(0)

	for _, payment := range payments {
		summPays += payment.Amount
	}

	return summPays / types.Money(len(payments))
}

// TotalInCategory находит сумму покупок в определенной категории.
func TotalInCategory(payments []types.Payment, category types.Category) types.Money{

	summPays := 0

	for _, payment := range payments {
		if payment.Category == category {
			summPays += int(payment.Amount)
		}
	}

	return types.Money(summPays)

}