package api

import (
	"Backend/cisco"
	"Backend/server"
	"fmt"
	"log"
)

func ChangeHostname (r *server.APIRequest) (err error) {
	hostname, ok := r.Request.Vars["hostname"]
	if !ok {
		log.Println("Url Param 'key' is missing")
		return
	}
	conn, err := cisco.Connect("192.168.1.5:22", "cisco", "class")
	if err != nil {
		fmt.Printf("There was an error connecting: %s", err)
	}
	_, err = conn.SendCommands([]string{"conf t", "hostname " + hostname})
	if err != nil {
		fmt.Printf("There was an error showing runnnig: %s", err)
	}
	return
}
