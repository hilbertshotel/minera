package handlers

var connStr string = "user=postgres dbname=minera_catalog sslmode=disable host=/run/postgresql"

type Category struct {
	Id int `json:id`
	Name string `json:name`
}

type Item struct {
	Name string `json:name`
	Description string `json:description`
	Images []string `json:images`
}
