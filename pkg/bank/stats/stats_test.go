package stats

import (
	"fmt"
	"reflect"
	"testing"
	"github.com/Jaborov-U/Day-11-BankTypes/v2/pkg/types"
)

func ExampleAVG() {

	cards := []types.Payment{
		{
			Amount: 10,
			Status: types.StatusFail,
		},
		{
			Amount: 15,
			Status: types.StatusInProgress,
		},
		{
			Amount: 20,
			Status: types.StatusOk,
		},
	}
	avgPays := AVG(cards)

	fmt.Println(avgPays)

	//Output:17

}

func ExampleTotalInCategory() {

	cards := []types.Payment{
		{
			Amount: 10,
			Category: "Shop",
			Status: types.StatusFail,
		},
		{
			Amount: 15,
			Category: "Shop",
			Status: types.StatusInProgress,
		},
		{
			Amount: 15,
			Category: "Internet",
			Status: types.StatusOk,
		},
	}
	avgPays := TotalInCategory(cards, "Shop")

	fmt.Println(avgPays)
	
	//Output: 15
	
}

func TestCategoriesAvg_Test(t *testing.T) {

	payments := []types.Payment{
		{Category: "internet", Amount: 10},
		{Category: "taxi", Amount: 10},
		{Category: "taxi", Amount: 20},
	}

	expected := map[types.Category]types.Money{
		"internet": 10,
		"taxi": 15,
	}

	result := CategoriesAvg(payments)

	if !reflect.DeepEqual(expected, result){
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}

}