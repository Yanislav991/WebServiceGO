package main

import (
	"net/http"
	"test/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3005", nil)
}
