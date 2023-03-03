package authorization

import (
	"booking/backend/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func RegistrationHandler(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	writer.WriteHeader(http.StatusOK)
	input := struct {
		UsernameOrEmail string `json:"username-or-email"`
		Password        string `json:"password"`
	}{}
	fmt.Println(json.NewDecoder(request.Body).Decode(&input))
	fmt.Println(input.UsernameOrEmail)
	fmt.Println(request.Method)
}

func LoginPage(writer http.ResponseWriter, request *http.Request) {
	err := utils.ExecuteTemplates(writer, nil, "templates/login.html")
	if err != nil {

	}
}
