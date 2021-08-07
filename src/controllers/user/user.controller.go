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
	r.HandleFunc(controllerName, create).Methods("DELETE")
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
	users, err := userService.GetAll()

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			config.NewHTTPError(err, http.StatusInternalServerError, "Error al cargar los clientes"))

		return
	}

	json.NewEncoder(w).Encode(users)

}
