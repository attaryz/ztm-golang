package main

import "fmt"

func main() {

	shoppingList := make(map[string]int)
	shoppingList["apples"] = 5
	shoppingList["oranges"] = 10
	shoppingList["bananas"] = 15
	shoppingList["bread"] += 3
	shoppingList["bananas"] += 1

	fmt.Println(shoppingList)

	delete(shoppingList, "apples")
	fmt.Println("apples got deleted", shoppingList)
	fmt.Println("need", shoppingList["oranges"], "oranges")

	cereal, found := shoppingList["cereal"]
	if found {
		fmt.Println("found", cereal)
	} else {
		fmt.Println("cereal not found")
	}

	milk, found := shoppingList["milk"]
	if !found {
		fmt.Println("milk not found shopping list has", shoppingList)
	} else {
		fmt.Println("milk found", milk)
	}

	totalItems := 0
	for _, amount := range shoppingList {
		totalItems += amount
	}
	fmt.Println("total items", totalItems)
}
