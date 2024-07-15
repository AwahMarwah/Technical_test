package model

type FruitType string

const (
	FruitTypeImport FruitType = "IMPORT"
	FruitTypeLocal  FruitType = "LOCAL"
)

type Fruit struct {
	ID    int       `json:"fruit_id"`
	Name  string    `json:"fruit_name"`
	Type  FruitType `json:"fruit_type"`
	Stock int       `json:"stock"`
}
