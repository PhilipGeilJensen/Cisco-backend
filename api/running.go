package api

import (
	"Backend/cisco"
	"Backend/resources"
	"encoding/json"
	"fmt"
	"net/http"
)


func ServeInterfaces(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Content-Type, YourOwnHeader")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
	}
	// Stop here if its Preflighted OPTIONS request
	if r.Method == "OPTIONS" {
		return
	}
	host, err := r.Cookie("host")
	username, err := r.Cookie("username")
	password, err := r.Cookie("password")
	if err != nil {
		fmt.Printf("There was an error getting the cookies.. %s", err)
	}

	c := resources.ConnectionCredentials{
		Host: host.Value,
		Username: username.Value,
		Password: password.Value,
	}

	lines, err := cisco.GetInterfaces(c)
	if err != nil {
		fmt.Printf("There was an error: %s", err)
	}

	err = json.NewEncoder(w).Encode(lines)
	if err != nil {
		fmt.Println("Shit")
	}
}