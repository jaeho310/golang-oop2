package infrastructure

import (
	"golang-oop/dto"
	"strconv"
)

type Kitchen struct{}

func (Kitchen) Cook(foodDto *dto.FoodDto) (string, error) {
	foodName := foodDto.Name
	foodCount := foodDto.Count

	return foodName + strconv.Itoa(foodCount) + " 개", nil
}
