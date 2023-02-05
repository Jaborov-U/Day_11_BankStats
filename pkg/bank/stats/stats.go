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

// PeriodsDynamic сравнивает расходы по категориям за два периода
func PeriodsDynamic(first map[types.Category]types.Money, second map[types.Category]types.Money) map[types.Category]types.Money{

	periodDinamic := map[types.Category]types.Money{}

	for key, value  := range second {
		_, ok := first[key]
		if !ok {
			periodDinamic[key] = 0
		} 
		periodDinamic[key] += value - first[key]
	}

	for key, value  := range first {
		_, ok := second[key]
		if !ok {
			periodDinamic[key] = 0 - value
		} 
	}
	
	return periodDinamic
}