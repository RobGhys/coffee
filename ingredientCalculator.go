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

func availableQ() {
	// Get quantities from user
	var qWater, qMilk, qBeans, qCupsNeeded int
	fmt.Println("Write how many ml of water the coffee machine has:")
	fmt.Scan(&qWater)
	fmt.Println("Write how many ml of milk the coffee machine has:")
	fmt.Scan(&qMilk)
	fmt.Println("Write how many grams of coffee beans the coffee machine has:")
	fmt.Scan(&qBeans)
	fmt.Println("Write how many cups of coffee you will need:")
	fmt.Scan(&qCupsNeeded)

	var nbrCupsAvailable int = computeNbCupsAvailable(qWater, qMilk, qBeans)
	//fmt.Printf("I can make %d cups", nbrCupsAvailable)

	fmt.Println(giveAnswer(qCupsNeeded, nbrCupsAvailable))
}

func giveAnswer(needed int, available int) string {
	if needed == available {
		return "Yes, I can make that amount of coffee"
	} else if needed > available {
		return fmt.Sprintf("%s %d %s", "No, I can make only ", available, " cups of coffee")
	} else {
		return fmt.Sprintf("%s %d %s", "Yes, I can make that amount of coffee (and even ", available-needed, " more than that)")
	}

}

func computeNbCupsAvailable(qWater int, qMilk int, qBeans int) int {
	var requiredQWater, requiredQMilk, requiredQBeans int = 200, 50, 15

	if qWater < requiredQWater || qMilk < requiredQMilk || qBeans < requiredQBeans {
		return 0
	} else {
		return 1 + computeNbCupsAvailable(qWater-requiredQWater, qMilk-requiredQMilk, qBeans-requiredQBeans)
	}
}
