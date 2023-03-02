package apiserver

import (
	"Go-http-rest-api/internal/api/store"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
	}
	s.configureRouter()
	return s
}
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request){
	s.router.ServeHTTP(w,r)
}
func (s *server) configureRouter() {
    s.router.HandleFunc("/users",s.handleUsersCteate()).Methods("POST")
}
func (s *server) handleUsersCteate() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request){
		
	}
}
