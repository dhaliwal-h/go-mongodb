package contorllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dhaliwal-h/go-mongodb/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userModel := models.User{}
	err := json.NewDecoder(r.Body).Decode(&userModel)
	if err != nil {
		fmt.Println(err)
	}
	userModel.Id = bson.NewObjectId()
	uc.session.DB("mongo-go").C("users").Insert(userModel)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(userModel)
}
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	fmt.Println(id)
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(id)
	u := models.User{}
	if err := uc.session.DB("mongo-go").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Println(err)
		return
	}

	uc.session.DB("mongo-go").C("users").RemoveId(oid)
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	fmt.Println(id)
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(id)
	u := models.User{}
	if err := uc.session.DB("mongo-go").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Println(err)
		return
	}

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)

}

func (uc UserController) GetAllUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var alluser []models.User

	uc.session.DB("mongo-go").C("users").Find(nil).All(&alluser)
	json.NewEncoder(w).Encode(alluser)
	fmt.Printf("%v", alluser)
}
