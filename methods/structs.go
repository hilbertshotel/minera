package methods

// CATEGORIES
type Category struct {
	Id int
	Name string
}

// SUB CATEGORIES
type SubCategory struct {
	Id int
	Name string
}

type SubTempData struct {
	CatId int
	CatName string
	Subs []SubCategory
}

// PRODUCTS
type Product struct {
	Id int `json:id`
	Name string `json:name`
	Description string `json:description`
	Images []string `json:images`
}

type ProductTempData struct {
	CatId int
	SubCatId int
	SubCatName string
	Products []Product
}