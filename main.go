package main

import (
	"fmt"
	"golang-oop/dto"
	"golang-oop/infrastructure"
)

func main() {
	myFood := &dto.FoodDto{Name: "스테이크", Count: 2}

	result, err := infrastructure.DiningHall{}.Order(myFood)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(result)
}
