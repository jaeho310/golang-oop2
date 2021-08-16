package infrastructure

import (
	"errors"
	"golang-oop/dto"
)

type DiningHall struct{}

func (diningHall DiningHall) Order(foodDto *dto.FoodDto) (string, error) {
	if foodDto.Name == "" || foodDto.Count <= 0 {
		return "", errors.New("주문을 확인해주세요")
	}
	response, err := Kitchen{}.Cook(foodDto)
	if err != nil {
		return "", err
	}
	return "주문하신 " + response + " 나왔습니다.", nil
}
