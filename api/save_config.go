package api

import (
	"Backend/cisco"
	"Backend/resources"
	"Backend/server"
	"net/http"
)

func SaveConfig(r *server.APIRequest) error {
	var config resources.ConnectionCredentials
	err := r.Decode(&config)
	if err != nil {
		http.Error(r.Request.W, "Error parsing the connection credentials", http.StatusNotAcceptable)
		return err
	}
	conn, err := cisco.Connect(config.Host + ":22", config.Username, config.Password)
	if err != nil {
		http.Error(r.Request.W, "Error getting a connection to the device", http.StatusNotAcceptable)
		return err
	}
	_, err = conn.SendShowCommand("wr")
	if err != nil {
		return err
	}
	return nil
}
