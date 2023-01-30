package stats

import (
	"fmt"
	"github.com/Jaborov-U/Day-11-BankTypes/pkg/types"
)

func ExampleAVG() {

	cards := []types.Payment{
		{
			Amount: 10,
		},
		{
			Amount: 15,
		},
		{
			Amount: 20,
		},
	}
	avgPays := AVG(cards)

	fmt.Println(avgPays)

	//Output:15

}

func ExampleTotalInCategory() {

	cards := []types.Payment{
		{
			Amount: 10,
			Category: "Shop",
		},
		{
			Amount: 10,
			Category: "Shop",
		},
		{
			Amount: 15,
			Category: "Internet",
		},
	}
	avgPays := TotalInCategory(cards, "Shop")

	fmt.Println(avgPays)
	
	//Output: 20
	
}
