package categories

import "net/http"

func Dispatcher(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
		case http.MethodGet: get(writer)
		case http.MethodPost: post(writer, request)
		case http.MethodPut: put(writer, request)
		case http.MethodDelete: delete(writer, request)
		default: http.Error(writer, "Възникна грешка", 405)
	}
}
