package stats

import "github.com/Jaborov-U/Day-11-BankTypes/v2/pkg/types"

// AVG рассчитывает среднюю сумму платежа.
func AVG(payments []types.Payment) types.Money {

	summPays := types.Money(0)

	count := types.Money(0)

	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			count++
			summPays += payment.Amount
		}
	}

	return summPays / count
}

// TotalInCategory находит сумму покупок в определенной категории.
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {

	summPays := 0

	for _, payment := range payments {
		if payment.Category == category && payment.Status != types.StatusFail {
			summPays += int(payment.Amount)
		}
	}

	return types.Money(summPays)

}

// CategoriesAvg считает среднюю сумму платежа по каждой категории.
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money{
    
    categories := map[types.Category]types.Money{}
    count := map[types.Category]int{}

    for _, payment := range payments {
        categories[payment.Category] += payment.Amount
        count[payment.Category] ++ 
    }

    for key, value := range count {
        categories[key] = categories[key] / types.Money(value)
    }

    return categories
}
