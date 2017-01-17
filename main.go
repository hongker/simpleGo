package main

import (
	"framework/controllers"
	"framework/vendor"
	"log"
	"net/http"
)

func main() {
	router := vendor.NewRouter()
	router.RegisterController("/", &controllers.IndexController{})
	router.RegisterController("/user", &controllers.UserController{})
	router.RegisterController("/user/hello", &controllers.UserController{})
	router.RegisterController("/user/info/:id", &controllers.UserController{})

	err := http.ListenAndServe(":8888", router)

	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
