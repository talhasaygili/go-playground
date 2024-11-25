package main

import (
	"fmt"
)


func main(){
	var costPrice, sellingPrice, profit float64
	fmt.Print("Enter Cost Price: ")
	fmt.Scan(&costPrice)

	fmt.Print("Enter Selling Price: ")
	fmt.Scan(&sellingPrice)

	profit = sellingPrice - costPrice

	if profit > 0 {
		fmt.Println("You made a profit of", profit)
	} else if profit < 0 {
		fmt.Println("You incurred a loss of", -profit)
	} else {
		fmt.Println("No Profit No Loss")
	}

}

