package models

type State string

const (
	Available State = "Available"
	Occupied  State = "Occupied"
	Pending   State = "Pending"
)

type Table struct {
	Name     string
	Capacity string
	State    State
}

func FetchTables() []Table {
	return []Table{
		{"#TBL007", "4", Available},
		{"#TBL013", "8", Occupied},
		{"#TBL009", "12", Pending},
	}
}
