package models

type Product struct {
	Name        string
	Description string
	Price       float64
	Image       string
}

func FetchProducts() []Product {
	return []Product{
		{
			"Luwombo Chicken",
			"Steamed chicken with a delicious taste, wrapped in banana leaves to keep the natural aroma.",
			35000.00,
			"../static/assets/imgs/lw-chicken.jpeg",
		},
		{
			"Spicy Rolex",
			"Prepared from the best wheat and vegetable oil plus eggs from locally bred poultry.",
			7000.00,
			"../static/assets/imgs/ff-rolex.jpeg",
		},
		{
			"Pineapple Juice",
			"Perfectly blended from organic fruits locally grown in Uganda with zero sugar added.",
			5000.00,
			"../static/assets/imgs/juice-pineapple.jpg",
		},
		{
			"Tropical Fruitsalad",
			"A mix of most tropical fruits, vegetables, berries, nuts and citrus fruits.",
			15000.00,
			"../static/assets/imgs/fr-fruit_salad.jpg",
		},
		{
			"Pilau & Goat",
			"Yummy brown rice with goat's meat. Not Biriyani, it's prepared in a local way.",
			5000.00,
			"../static/assets/imgs/st-pilau.jpeg",
		},
		{
			"Ettooke Eriboobedde",
			"Steamed matooke/bananas wrapped in banana leaves to keep the natural aroma.",
			5000.00,
			"../static/assets/imgs/st-matooke.jpeg",
		},
	}
}
