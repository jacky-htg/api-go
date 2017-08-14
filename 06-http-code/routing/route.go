package routing

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/jacky-htg/api-go/06-http-code/controllers"
)

type Route struct {
	Path string
	Method string
	Handler http.HandlerFunc
}

type Routes []Route

var routes = Routes {
	Route {"/api/get-token","POST",controllers.GetTokenHandler },
	Route {"/api/users","GET",controllers.GetUsersEndPoint },
	Route {"/api/users","POST",controllers.CreateUserEndPoint },
	Route {"/api/users/{id}","GET",controllers.GetUserEndPoint },
	Route {"/api/users/{id}","PUT",controllers.EditUserEndPoint },
	Route {"/api/users/{id}","DELETE",controllers.DeleteUserEndPoint },
}


func NewRouter() (*mux.Router) {
	router := mux.NewRouter()

	for _, route := range routes {
		router.HandleFunc(route.Path,  route.Handler).Methods(route.Method)
	}

	return router
}
