package backend

import "html/template"

const Address = "127.0.0.1:8000"
var templates = template.Must(template.ParseGlob("templates/*"))


type Category struct {
	Name string
	Path string
}

type Item struct {
	Name string
	Description string
	Images []string
}

var function_map = map[string]interface{}{
	"list_items": list_items,
}
