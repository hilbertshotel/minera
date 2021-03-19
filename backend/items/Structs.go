package items

type Item struct {
	Id int `json:id`
	Name string `json:name`
	Description string `json:description`
	Images []string `json:images`
}
