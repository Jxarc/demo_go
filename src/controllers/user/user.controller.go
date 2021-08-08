package usercontroller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	config "github.com/axel526/jikkosoft/src/controllers/config"
	userService "github.com/axel526/jikkosoft/src/services/user.service"
)

//InitController inicializa todos los endpoint para este controller
func InitController(controllerName string, r *mux.Router) {
	r.HandleFunc(controllerName, getAll).Methods("GET")
	r.HandleFunc(controllerName+"/{userId}", getUserDetails).Methods("GET")
	r.HandleFunc(controllerName, delete).Methods("DELETE")
	r.HandleFunc(controllerName, update).Methods("PUT")
	r.HandleFunc(controllerName, create).Methods("POST")
}

func create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create")
}

func delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "delete")
}

func update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "update")
}

func getAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := userService.GetAll()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			config.NewHTTPError(err, http.StatusInternalServerError, "Error al cargar usuarios"))

		return
	}

	json.NewEncoder(w).Encode(users)

}

func getUserDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	userId := vars["userId"]

	user, err := userService.GetUser(userId)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			config.NewHTTPError(err, http.StatusInternalServerError, "Ocurrío un problema al otener la data"+err.Error()))
		return
	}

	json.NewEncoder(w).Encode(user)
}
