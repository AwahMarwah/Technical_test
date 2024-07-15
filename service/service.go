package services

import (
	"github.com/AwahMarwah/Technical_test/model"
)

type FruitService struct {
	fruits []model.Fruit
}

func NewFruitService() *FruitService {
	return &FruitService{
		fruits: []model.Fruit{
			{ID: 1, Name: "Apel", Type: model.FruitTypeImport, Stock: 10},
			{ID: 2, Name: "Kurma", Type: model.FruitTypeImport, Stock: 20},
			{ID: 3, Name: "apel", Type: model.FruitTypeImport, Stock: 50},
			{ID: 4, Name: "Manggis", Type: model.FruitTypeLocal, Stock: 100},
			{ID: 5, Name: "Jeruk Bali", Type: model.FruitTypeLocal, Stock: 10},
			{ID: 6, Name: "KURMA", Type: model.FruitTypeImport, Stock: 20},
			{ID: 7, Name: "Salak", Type: model.FruitTypeLocal, Stock: 150},
		},
	}
}

func (s *FruitService) GetAllFruitNames() []string {
	fruitNames := make([]string, len(s.fruits))
	for i, fruit := range s.fruits {
		fruitNames[i] = fruit.Name
	}
	return fruitNames
}

func (s *FruitService) SeparateFruitsByType() map[model.FruitType][]model.Fruit {
	separatedFruits := make(map[model.FruitType][]model.Fruit)
	for _, fruit := range s.fruits {
		separatedFruits[fruit.Type] = append(separatedFruits[fruit.Type], fruit)
	}
	return separatedFruits
}

func (s *FruitService) GetTotalStockByType() map[model.FruitType]int {
	totalStock := make(map[model.FruitType]int)
	for _, fruit := range s.fruits {
		totalStock[fruit.Type] += fruit.Stock
	}
	return totalStock
}

func (s *FruitService) Comments() string {
	return "Komentar: Dalam data yang diberikan, kita dapat melihat bahwa ada buah dengan nama yang sama tetapi dengan perbedaan dalam huruf kapital, yaitu 'Apel' dan 'apel'. Hal ini bisa menyebabkan kebingungan dan ketidakakuratan dalam pengelolaan data. Nama buah 'Apel' dan 'apel' dianggap sebagai dua entitas yang berbeda karena perbedaan dalam huruf kapital. Dalam sistem yang case-sensitive, ini akan dianggap sebagai dua buah yang berbeda."
}
