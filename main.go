package main

import (
	"fmt"
	"golang-oop/dto"
	"golang-oop/infrastructure"
)

func main() {
	myFood := dto.FoodDto{}.New("스테이크", 2)

	kitChen := infrastructure.Kitchen{}.New()
	diningHall := infrastructure.DiningHall{}.New(kitChen)

	result, err := diningHall.Order(myFood)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(result)
}
