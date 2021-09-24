package api

import (
	"Backend/cisco"
	"Backend/resources"
	"Backend/server"
	"net/http"
)

func ConfigureInterface(r *server.APIRequest) error {
	var config resources.ConfigureInterface
	err := r.Decode(&config)
	if err != nil {
		http.Error(r.Request.W, "There was an error parsing the body", http.StatusBadRequest)
	}
	err = cisco.ConfigureInterface(config)
	if err != nil {
		http.Error(r.Request.W, "There was an error executing the command", http.StatusBadRequest)
	}
	return nil
}
