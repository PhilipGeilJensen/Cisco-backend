package api

import (
	"Backend/cisco"
	"Backend/resources"
	"Backend/server"
	"net/http"
)

func GetSnmpInfo(r *server.APIRequest) error {

	var c resources.SnmpCredentials

	err := r.Decode(&c)
	if err != nil {
		http.Error(r.Request.W, "There was an error parsing the body", http.StatusBadRequest)
	}

	r.Encode(cisco.Snmp(c))
	return nil
}
