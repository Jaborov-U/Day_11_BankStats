package stats

import (
	"fmt"
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
