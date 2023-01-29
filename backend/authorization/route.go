package authorization

import (
	"fmt"
	"net/http"
)

func RegistrationHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "HelloWorld")
}
