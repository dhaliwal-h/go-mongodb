package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dhaliwal-h/go-mongodb/contorllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := contorllers.NewUserController(getSession())

	r.HandlerFunc("GET", "/", handleHome)
	r.GET("/user/:id", uc.GetUser)
	r.GET("/users", uc.GetAllUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	fmt.Println("Stariting server at port 8090")
	log.Fatal(http.ListenAndServe(":8090", r))

}

func handleHome(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("helloworld")
}
func getSession() *mgo.Session {
	se, err := mgo.Dial("mongodb://127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	return se
}
