package items

import (
	"net/http"
)

func Dispatcher(writer http.ResponseWriter, request *http.Request, id int) {
	switch request.Method {
		case http.MethodGet: Get(writer, id)
		case http.MethodPost: Post(writer, request)
		case http.MethodPut: Put(writer, request)
		case http.MethodDelete: Delete(writer, id)
		default: return
	}
}