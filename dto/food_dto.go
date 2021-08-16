package dto

type FoodDto struct {
	name  string
	count int
}

func (FoodDto) New(name string, count int) *FoodDto {
	return &FoodDto{name: name, count: count}
}

func (food *FoodDto) SetName(name string) {
	food.name = name
}

func (food *FoodDto) GetName() string {
	return food.name
}

func (food *FoodDto) SetCount(count int) {
	food.count = count
}

func (food *FoodDto) GetCount() int {
	return food.count
}
