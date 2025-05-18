package models

type BoughtItem struct {
	ProductID   string
	ProductName string
	Quantity    int
	Price       float64
}

type Order struct {
	ID         string
	TableID    string
	Items      []BoughtItem
	Status     string
	CreatedAt  string
	TotalPrice float64
}
