package categories

import (
	"net/http"
)

func Dispatcher(writer http.ResponseWriter, request *http.Request) {
	// METHOD MATCHING
	switch request.Method {
		case http.MethodGet: Get(writer)
		case http.MethodPost: Post(writer, request)
		case http.MethodPut: Put(writer, request)
		case http.MethodDelete: Delete(writer, request)
		default: return
	}

}