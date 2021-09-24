package server

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

//APIRequest is given as a parameter to every API function
type APIRequest struct {
	Request *Request
}

//APIHandler function for setting the template of the desired API functions
type APIHandler func(r *APIRequest) error

//API function implements the functions, and determines the methods, path and APIHandler method
func (srv *Server) API(methods, path string, handler APIHandler) {
	srv.mux.Handle(path, &metricsHandler{path, &apiHandler{
		handler: handler,
		srv:     srv,
	}}).Methods(strings.Split(methods, ",")...)
}

//apiHandler is a private struct which holds the public APIHandler and a pointer to the *Server instance
type apiHandler struct {
	handler APIHandler
	srv     *Server
}

//ServeHTTP function of the apiHandler
func (h *apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	req := &APIRequest{
		Request: request(w, r, h.srv),
	}
	err := h.handler(req)
	if err != nil {
		logrus.Errorf("API handler returned error: %s", err)
	}
}

//Decode is a function for decoding the HTTP body from JSON to a interface (in this project, just structs)
func (r *APIRequest) Decode(p interface{}) error {
	jsonEncoded := strings.Contains(r.Request.R.Header.Get("Content-Type"), "application/json")
	if !jsonEncoded {
		return nil
	}
	err := json.NewDecoder(r.Request.R.Body).Decode(p)
	if err != nil {
		logrus.Warnf("Decode JSON error: %s", err)
		return err
	}
	return nil
}
//Encodes the provided interface into JSON, and sends the formatted JSON to the http.ResponseWriter
func (r *APIRequest) Encode(p interface{}) bool {
	r.Request.R.Header.Set("Content-Type", "application/json; charset=utf-8")
	err := json.NewEncoder(r.Request.W).Encode(p)
	if err != nil {
		return false
	}
	return true
}
