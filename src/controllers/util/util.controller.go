package utilcontroller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	config "github.com/axel526/jikkosoft/src/controllers/config"
	e "github.com/axel526/jikkosoft/src/entity"
	utilService "github.com/axel526/jikkosoft/src/services/util.service"
	"github.com/gorilla/mux"
)

//InitController inicializa todos los endpoint para este controller
func InitController(controllerName string, r *mux.Router) {
	r.HandleFunc(controllerName+"/sort", sortedArray).Methods("POST")
}

func sortedArray(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	var sorted e.Sorted
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			config.NewHTTPError(err, http.StatusInternalServerError, err.Error()))
	}

	json.Unmarshal(reqBody, &sorted)

	if sorted.Unsorted == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			config.NewHTTPError(err, http.StatusBadRequest, "Error al procesar input"))

	} else {
		json.NewEncoder(w).Encode(utilService.Sorted(sorted.Unsorted))
	}

}
