package categories

const connStr string = "user=postgres dbname=minera_catalog sslmode=disable host=/run/postgresql"
// const connStr string = "user=postgres dbname=minera_catalog sslmode=disable" // WINDOWS

type Category struct {
	Id int `json:id`
	Name string `json:name`
}

type PutCategory struct {
	Id int `json:id`
	NewName string `json:newName`
}