package models

type Stock string

const (
	True  Stock = "True"
	False Stock = "False"
)

type Product struct {
	Name        string
	Description string
	Price       string
	AddedDate   string
	Image       string
	Stock       Stock
}

func FetchProducts() []Product {
	return []Product{
		{"Luwombo Chicken",
			"Steamed chicken with a delicious taste, wrapped in banana leaves to keep the natural aroma.",
			"35000",
			"13-05-2025",
			"../static/assets/",
			"True",
		},
		{"Spicy Rolex",
			"Prepared from the best wheat and vegetable oil plus eggs from locally bred poultry.",
			"7000",
			"13-05-2025",
			"../static/assets/",
			"True",
		},
		{"Pineapple Juice",
			"Perfectly blended from organic fruits locally grown in Uganda with zero sugar added.",
			"5000",
			"13-05-2025",
			"../static/assets/",
			"True",
		},
		{"Tropical Fruitsalad",
			"A mix of most tropical fruits, vegetables, berries, nuts and citrus fruits.",
			"15000",
			"13-05-2025",
			"../static/assets/",
			"True",
		},
		{"Pilau & Goat",
			"Yummy brown rice with goat's meat. Not Biriyani, it's prepared in a local way.",
			"5000",
			"13-05-2025",
			"../static/assets/",
			"True",
		},
		{"Ettooke Eriboobedde",
			"Steamed matooke/bananas wrapped in banana leaves to keep the natural aroma.",
			"7.57",
			"13-05-2025",
			"../static/assets/",
			"True",
		},
	}
}
