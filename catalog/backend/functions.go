package backend

import (
	"fmt"
)


// MISC
func search_path(path string) (string, string, int, bool) {
	// SELECT function_name, section, id WHERE path = path;
	return "list_items", "sub", 7, true
}


// CATEGORIES
func get_categories() []Category {
	categories := []Category{
		Category{ Name: "Category 1", Path: "http://" + Address + "/345" },
		Category{ Name: "Category 2", Path: "http://" + Address + "/127" },
		Category{ Name: "Category 3", Path: "http://" + Address + "/325" },
	}
	return categories
}

// SUB CATEGORIES
func list_items() {
	fmt.Println("list items works")
}