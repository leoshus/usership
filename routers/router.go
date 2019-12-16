/* Package provider request route handler */
package routers

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"usership/config"
	"usership/controllers"
)

type Service struct {
	Router *mux.Router
}

//database and router initialize
func (s *Service) InitAndRun() {
	s.Router = mux.NewRouter()
	s.initRouters()
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", config.Get().Web.Host, fmt.Sprint(config.Get().Web.Port)), s.Router))
}

//register handler with url to router
func (s *Service) initRouters() {
	s.Router.HandleFunc("/users", controllers.ListUser).Methods("GET")
	s.Router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	s.Router.HandleFunc("/users/{user_id}/relationships", controllers.QueryRelationShip).Methods("GET")
	s.Router.HandleFunc("/users/{user_id}/relationships/{other_user_id}", controllers.ChangeState).Methods("PUT")
}
