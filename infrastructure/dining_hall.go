package infrastructure

import (
	"errors"
	"golang-oop/dto"
)

type DiningHall struct {
	kitchen *Kitchen
}

func (DiningHall) New(kitchen *Kitchen) *DiningHall {
	return &DiningHall{kitchen}
}

func (diningHall *DiningHall) Order(foodDto *dto.FoodDto) (string, error) {
	if foodDto.GetName() == "" || foodDto.GetCount() <= 0 {
		return "", errors.New("주문을 확인해주세요")
	}
	response, err := diningHall.kitchen.Cook(foodDto)
	if err != nil {
		return "", err
	}
	return "주문하신 " + response + " 나왔습니다.", nil
}
