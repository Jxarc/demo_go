package api

import (
	"fmt"
	"net/http"

	usercontroller "github.com/axel526/jikkosoft/src/controllers/user"
	"github.com/gorilla/mux"
)

//InitAPI inicializar todos los controladores del api
func InitAPI(port string) {
	router := mux.NewRouter().StrictSlash(true)

	usercontroller.InitController("/users", router)

	fmt.Println("API is running")
	http.ListenAndServe(":"+port, router)
}
