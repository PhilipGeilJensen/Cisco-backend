package cisco

import (
	"Backend/resources"
	"fmt"
)

func GetBannerMotd(connection resources.ConnectionCredentials) (banner string, err error) {
	conn, err := Connect(connection.Host + ":22", connection.Username, connection.Password)
	if err != nil {
		fmt.Printf("There was an error connecting: %s", err)
		return
	}
	run, err := conn.SendShowCommand("show banner motd")
	if err != nil {
		fmt.Printf("There was an error showing runnnig")
		return
	}
	banner = string(run)
	return
}
