package main

import (
	"database/sql"
	"jwt-course-refactored/controllers"

	"jwt-course-refactored/driver"
	"jwt-course-refactored/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var db *sql.DB

func init() {
	gotenv.Load()
}

func main() {
	db = driver.ConnectDB()
	router := mux.NewRouter()

	controller := controllers.Controller{}

	router.HandleFunc("/protectedEndpoint", utils.TokenVerifyMiddleWare(controller.Protected(db))).Methods("GET")
	router.HandleFunc("/signup", controller.Signup(db)).Methods("POST")
	router.HandleFunc("/login", controller.Login(db)).Methods("POST")
	router.HandleFunc("/test", test).Methods("GET")

	log.Println(http.ListenAndServe(":8000", router))
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("succ"))
}
