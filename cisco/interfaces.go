package cisco

import (
	"Backend/resources"
	"fmt"
)

func GetInterfaces(connection resources.ConnectionCredentials) (lines []resources.Interface, err error) {
	conn, err := Connect(connection.Host + ":22", connection.Username, connection.Password)
	if err != nil {
		fmt.Printf("There was an error connecting: %s", err)
		return
	}
	run, err := conn.SendShowCommand("show ip interface brief")
	if err != nil {
		fmt.Printf("There was an error showing runnnig")
		return
	}
	lines = resources.FormatInterface(run)
	return
}

func ConfigureInterface(configuration resources.ConfigureInterface) error {
	conn, err := Connect(configuration.Host + ":22", configuration.Username, configuration.Password)
	if err != nil {
		fmt.Printf("There was an error connecting: %s", err)
		return err
	}
	var shut string
	if configuration.Shutdown {
		shut = "shutdown"
	} else {
		shut = "no shutdown"
	}
	_, err = conn.SendCommands([]string{"conf t", "interface " + configuration.Interface, shut, "ip address " + configuration.IpAddress + " " + configuration.SubnetMask})
	if err != nil {
		fmt.Printf("There was an error showing runnnig")
		return err
	}
	return nil
}