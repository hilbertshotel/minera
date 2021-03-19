package items

import (
	"net/http"
)

func Dispatcher(writer http.ResponseWriter, request *http.Request, id int) {
	// METHOD MATCHING
	switch request.Method {
		case http.MethodGet: Get(writer, request, id)
		case http.MethodPost: Post(writer, request)
		case http.MethodPut: Put(writer, request)
		case http.MethodDelete: Delete(writer, request)
		default: return
	}

}