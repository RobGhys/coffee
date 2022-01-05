package main

import "fmt"

func start(qWater *int, qMilk *int, qBeans *int, qCups *int, money *int) {
	showRemainingQ(qWater, qMilk, qBeans, qCups, money)
	doAction(qWater, qMilk, qBeans, qCups, money)
}

func doAction(qWater *int, qMilk *int, qBeans *int, qCups *int, money *int) {
	// Prompt the user
	var userAction string

	fmt.Println("Write action (buy, fill, take):")
	fmt.Scan(&userAction)

	switch userAction {
	case "buy":
		buy(qWater, qMilk, qBeans, qCups, money)
	case "fill":
		fill(qWater, qMilk, qBeans, qCups, money)
	case "take":
		take(qWater, qMilk, qBeans, qCups, money)
	default:
		fmt.Println("I don't understand what you are asking me.")
	}
}

func take(qWater *int, qMilk *int, qBeans *int, qCups *int, money *int) {
	fmt.Printf("I give you $ %d\n", *money)
	*money = 0
	showRemainingQ(qWater, qMilk, qBeans, qCups, money)
}

func fill(qWater *int, qMilk *int, qBeans *int, qCups *int, money *int) {
	var addWater, addMilk, addBeans, addCups int

	fmt.Println("Write how many ml of water you want to add:")
	fmt.Scan(&addWater)
	fmt.Println("Before: ", *qWater)
	*qWater += addWater
	fmt.Println("After:", *qWater)

	fmt.Println("Write how many ml of milk you want to add:")
	fmt.Scan(&addMilk)
	*qMilk += addMilk

	fmt.Println("Write how many grams of coffee beans you want to add:")
	fmt.Scan(&addBeans)
	*qBeans += addBeans

	fmt.Println("Write how disposable coffee cups you want to add:")
	fmt.Scan(&addCups)
	*qCups += addCups

	showRemainingQ(qWater, qMilk, qBeans, qCups, money)
}

func buy(qWater *int, qMilk *int, qBeans *int, qCups *int, money *int) {
	var choice int

	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino:")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		var reqWater, reqMilk, reqBeans, reqCup, reqMoney int = 250, 0, 16, 1, 4
		buyDrink(qWater, qMilk, qBeans, qCups, money, reqWater, reqMilk, reqBeans, reqCup, reqMoney)
	case 2:
		var reqWater, reqMilk, reqBeans, reqCup, reqMoney int = 350, 75, 20, 1, 7
		buyDrink(qWater, qMilk, qBeans, qCups, money, reqWater, reqMilk, reqBeans, reqCup, reqMoney)
	case 3:
		var reqWater, reqMilk, reqBeans, reqCup, reqMoney int = 200, 100, 12, 1, 6
		buyDrink(qWater, qMilk, qBeans, qCups, money, reqWater, reqMilk, reqBeans, reqCup, reqMoney)
	default:
		fmt.Println("I don't understand what you are asking me.")
	}
}

func buyDrink(qWater *int, qMilk *int, qBeans *int, qCups *int, money *int,
	reqWater int, reqMilk int, reqBeans int, reqCup int, reqMoney int) {

	// Decrease all amounts by the cost in ingredient of the drink
	*qWater -= reqWater
	*qMilk -= reqMilk
	*qBeans -= reqBeans
	*qCups -= reqCup
	*money += reqMoney

	showRemainingQ(qWater, qMilk, qBeans, qCups, money)
}

func showRemainingQ(qWater *int, qMilk *int, qBeans *int, qCups *int, money *int) {

	fmt.Println("The coffee machine has:")
	fmt.Println(*qWater, " of water")
	fmt.Println(*qMilk, " of milk")
	fmt.Println(*qBeans, " of coffee beans")
	fmt.Println(*qCups, " of disposable cups")
	fmt.Println(*money, " of money")
}
