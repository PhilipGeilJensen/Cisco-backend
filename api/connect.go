package api

import (
	"Backend/cisco"
	"Backend/resources"
	"Backend/server"
	"fmt"
	"net/http"
)

func CreateConnection(r *server.APIRequest) error {
	var c resources.ConnectionCredentials

	err := r.Decode(&c)
	if err != nil {
		http.Error(r.Request.W, "There was an error with the body", http.StatusNotAcceptable)
	}

	lines, err := cisco.GetInterfaces(c)
	if err != nil {
		fmt.Printf("There was an error: %s", err)
	}

	vlans, err := cisco.GetVlans(c)
	if err != nil {
		fmt.Printf("There was an error getting the vlans: %s", err)
	}

	banner, err := cisco.GetBannerMotd(c)
	if err != nil {
		fmt.Printf("There was an error getting the banner: %s", err)
	}

	conf := resources.Config{
		Interfaces: lines,
		Vlans:      vlans,
		Banner:     banner,
	}

	ok := r.Encode(&conf)
	if !ok {
		fmt.Println("Shit")
	}
	return nil
}
