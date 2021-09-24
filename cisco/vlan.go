package cisco

import (
	"Backend/resources"
	"fmt"
)

func GetVlans(connection resources.ConnectionCredentials) (vlans []resources.Vlan, err error){
	conn, err := Connect(connection.Host + ":22", connection.Username, connection.Password)
	if err != nil {
		fmt.Printf("There was an error connecting: %s", err)
		return
	}
	run, err := conn.SendShowCommand("show vlan brief")
	if err != nil {
		fmt.Printf("There was an error showing runnnig")
		return
	}
	vlans = resources.FormatVlan(run)
	return
}
