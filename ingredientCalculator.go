package main

import "fmt"

func askQuestion() {
	fmt.Println("Write how many cups of coffee you will need:")
	var nbCups int
	fmt.Scan(&nbCups)

	showQuantities(nbCups)
}

func computeQWater(nbCups int) int {
	return nbCups * 200
}

func computeQMilk(nbCups int) int {
	return nbCups * 50
}

func computeQCoffeeBeans(nbCups int) int {
	return nbCups * 15
}

func showQuantities(nbCups int) {
	var qWater int = computeQWater(nbCups)
	var qMilk int = computeQMilk(nbCups)
	var qCoffeeBeans int = computeQCoffeeBeans(nbCups)

	fmt.Println("For", nbCups, "cups of coffee you will need")
	fmt.Println(qWater, " ml of water")
	fmt.Println(qMilk, " ml of milk")
	fmt.Println(qCoffeeBeans, " g of coffee beans")
}
