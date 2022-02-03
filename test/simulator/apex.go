package simulator

import (
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

// NewController returns a new Apex Controller simulator.
func NewController(login, password string) *Controller {
	ctrl := &Controller{
		router:   mux.NewRouter(),
		sids:     make(map[string]struct{}),
		login:    login,
		password: password,
	}

	ctrl.router.HandleFunc("/rest/login", ctrl.handlePostLogin).Methods("POST")
	ctrl.router.HandleFunc("/rest/user", ctrl.handleGetUser).Methods("GET")

	return ctrl
}

// Controller simulates a local Apex controller.
type Controller struct {
	router   *mux.Router
	mutex    sync.RWMutex
	login    string
	password string
	sids     map[string]struct{}
}

// ServeHTTP implements http.Handler.
func (s *Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
