package routers

import (
	"github.com/gorilla/mux"
	"github.com/shijuvar/go-web/taskmanager/controllers"
)

//SetUserRoutes function receives a pointer to the Gorilla mux router object (mux.Router) as an
//argument and returns the pointer of the mux.Router object
func SetUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/users/register", controllers.Register).Methods("POST")
	router.HandleFunc("/users/login", controllers.Login).Methods("POST")
	return router
}
