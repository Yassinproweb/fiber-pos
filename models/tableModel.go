package models

type State string

const (
	Available State = "Available"
	Occupied  State = "Occupied"
	Pending   State = "Pending"
)

type Table struct {
	Name      string
	Capacity  int
	State     State
	OrderCurr *Order
}

func FetchTables() []Table {
	return []Table{
		{
			Name:      "#TBR007",
			Capacity:  6,
			State:     Available,
			OrderCurr: nil,
		},
	}
}
