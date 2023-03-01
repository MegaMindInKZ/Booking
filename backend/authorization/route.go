package authorization

import (
	"booking/backend/utils"
	"fmt"
	"net/http"
)

func RegistrationHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello World")
}

func LoginPage(writer http.ResponseWriter, request *http.Request) {
	err := utils.ExecuteTemplates(writer, nil, "templates/login.html")
	if err != nil {

	}
}
