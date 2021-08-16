package infrastructure

import (
	"golang-oop/dto"
	"strconv"
)

type Kitchen struct{}

func (Kitchen) New() *Kitchen {
	return &Kitchen{}
}

func (*Kitchen) Cook(foodDto *dto.FoodDto) (string, error) {
	foodName := foodDto.GetName()
	foodCount := foodDto.GetCount()

	return foodName + strconv.Itoa(foodCount) + " 개", nil
}
