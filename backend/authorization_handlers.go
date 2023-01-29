package main

import "booking/backend/authorization"

func setAuthorizationHandlers() {
	Router.HandleFunc("/api/registration", authorization.RegistrationHandler)
}
