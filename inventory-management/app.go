package main

import (
	"fmt"
	"strings"
)

var ItemName string
var quantity int

var inventory = make(map[string]int)

func main() {
	var choice int


	fmt.Println("Welcome to the Inventory Management System")
	for choice != 4 {
	ShowMenu()
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)
	switch choice {
		case 1:
			AddItem()
		case 2:
			RemoveItem()
		case 3:
			ViewInventory()
		case 4:
			fmt.Println("Exiting the program")
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func ShowMenu(){
	fmt.Println("\n1. Add Item")
	fmt.Println("2. Remove Item")
	fmt.Println("3. View Inventory")
	fmt.Println("4. Exit")
}

func AddItem(){
	fmt.Print("Enter item name: ")
	fmt.Scan(&ItemName)

	fmt.Print("Enter quantity: ")
	fmt.Scan(&quantity)

	if val, ok := inventory[ItemName]; ok {
		inventory[ItemName] = val + quantity
	} else {
		inventory[ItemName] = quantity
	}
}

func ViewInventory(){
    if len(inventory) == 0 {
        fmt.Println("\nThe inventory is empty.")
        return
    }
    fmt.Print("\nInventory\n")
    fmt.Printf("%-15s %-10s\n", "Item", "Quantity")
    fmt.Println(strings.Repeat("-", 25))
    for itemName, quantity := range inventory {
        fmt.Printf("%-15s %-10d\n", itemName, quantity)
    }
}


func RemoveItem(){
	fmt.Print("\nEnter item name: ")
	fmt.Scan(&ItemName)

	fmt.Print("\nEnter quantity: ")
	fmt.Scan(&quantity)

	if val, ok := inventory[ItemName]; ok{
		if val >= quantity {
			remainedQuantity := val - quantity
			if remainedQuantity == 0 {
				delete(inventory, ItemName)
				return
			}
			inventory[ItemName] = val - quantity
		}else {
			fmt.Println("Insufficient quantity")
		}
	}
}