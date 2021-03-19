package backend

import (
	"strconv"
	"strings"
	"net/http"
	"minera/backend/categories"
	"minera/backend/items"
	"minera/backend/utils"
)

func Editor(writer http.ResponseWriter, request *http.Request) {
	
	// VALIDATOR
	cookie, err := request.Cookie(utils.CookieName)
	if err != nil {	utils.Logger.Println(err); return }

	if cookies[cookie.Name] != cookie.Value  {
		utils.Logger.Println("no such cookie `" + cookie.Name + "` in cookie jar")
		return
	}

	// URL PARSER
	u := request.URL.String()
	uArray := strings.Split(u, "/")
	section := uArray[2]

	// LOCAL DISPATCHER
	switch section {
		case "Categories":
			categories.Dispatcher(writer, request)
		
		case "Items":
			id, err := strconv.Atoi(uArray[3])
			if err != nil { utils.Logger.Println(err); return }
			items.Dispatcher(writer, request, id)

		default: return
	}

}
