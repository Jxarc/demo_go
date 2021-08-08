package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	usercontroller "github.com/axel526/jikkosoft/src/controllers/user"
	utilcontroller "github.com/axel526/jikkosoft/src/controllers/util"
)

//InitAPI inicializar todos los controladores del api
func InitAPI(port string) {
	router := mux.NewRouter().StrictSlash(true)

	usercontroller.InitController("/users", router)
	utilcontroller.InitController("/utils", router)

	fmt.Println("API is running")
	http.ListenAndServe(":"+port, router)
}
