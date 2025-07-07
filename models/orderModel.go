package models

type Type string
type Status string

const (
	Takeaway Type = "Takeaway"
	Delivery Type = "Delivery"
	DineIn   Type = "DineIn"

	Placed    Status = "Placed"
	Preparing Status = "Preparing"
	Ready     Status = "Ready"
	Canceled  Status = "Canceled"
	Transit   Status = "Transit"
	Delivered Status = "Delivered"
	Taken     Status = "Taken"
	Served    Status = "Served"
)

type Order struct {
	Name        string
	Type        Type
	Status      Status
	Items       string
	Cost        string
	CustName    string
	CustNumber  string
	Destination string
	DateTime    string
}

func FetchOrders() []Order {
	return []Order{
		{"#ORD0011", DineIn, Placed, "2", "7.33", "Ahmad", "0722678837", "Nakasozi, Wakiso", "11-06-2025_09:30"},
		{"#ORD0012", Takeaway, Taken, "1", "3.85", "Kasagga", "0767883721", "Kampala Branch", "11-06-2025_10:30"},
		{"#ORD0013", DineIn, Canceled, "7", "22.00", "Farīdah", "0762678030", "#TBL007", "11-06-2025_11:30"},
		{"#ORD0014", Delivery, Transit, "8", "102.79", "Josephine", "0727658937", "Wakaliga, Lubaga", "11-06-2025_12:30"},
		{"#ORD0015", DineIn, Served, "5", "33.5", "Sharīfah", "0742990939", "#TBL011", "11-06-2025_13:30"},
		{"#ORD0016", Delivery, Delivered, "13", "123.84", "Brian", "0700678111", "Kitende, Entebbe", "11-06-2025_14:30"},
		{"#ORD0017", Takeaway, Ready, "2", "13.75", "Rugaaju", "0755673337", "Mbarara Branch", "11-06-2025_15:30"},
		{"#ORD0018", DineIn, Preparing, "3", "23.05", "Okolot", "0779508837", "#TBL003", "11-06-2025_16:30"},
	}
}
