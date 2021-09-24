package server

import (
	"net/http"

	gmux "github.com/gorilla/mux"
)

// Request struct for holding usefull properties in the API functions
type Request struct {
	W http.ResponseWriter
	R *http.Request

	srv *Server

	Vars   map[string]string //URL parameters - (http://localhost:8080/api/host/{id})
	Method string //Which method has been used? POST, GET, PUT, OPTIONS, DELETE etc..
}

//Creates a new Request
func request(w http.ResponseWriter, r *http.Request, srv *Server) *Request {
	req := new(Request)
	req.W = w
	req.R = r
	req.Vars = gmux.Vars(r)
	req.Method = r.Method
	req.srv = srv
	return req
}
