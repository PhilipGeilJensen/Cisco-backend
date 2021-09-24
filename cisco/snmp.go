package cisco

import (
	"Backend/resources"
	"Backend/server"
	"fmt"
	g "github.com/gosnmp/gosnmp"
	"log"
	"os"
	"time"
)

func Snmp(credentials resources.SnmpCredentials) resources.SnmpInfo {

	params := &g.GoSNMP{
		Target:        credentials.Host,
		Port:          161,
		Version:       g.Version3,
		SecurityModel: g.UserSecurityModel,
		MsgFlags:      g.AuthPriv,
		Timeout:       time.Duration(30) * time.Second,
		SecurityParameters: &g.UsmSecurityParameters{UserName: credentials.User,
			AuthenticationProtocol:   g.MD5,
			AuthenticationPassphrase: credentials.Authentication,
			PrivacyProtocol:          g.AES,
			PrivacyPassphrase:        credentials.Privacy,
		},
	}
	err := params.Connect()
	if err != nil {
		log.Printf("Connect() err: %v", err)
	}
	defer params.Conn.Close()

	oids := []string{"1.3.6.1.2.1.1.1.0", "1.3.6.1.2.1.1.3.0", "1.3.6.1.4.1.9.2.1.8.0"}
	result, err2 := params.Get(oids) // Get() accepts up to g.MAX_OIDS
	if err2 != nil {
		log.Printf("Get() err: %v", err2)
	}

	return resources.SnmpInfo{
		SysInfo:    string(result.Variables[0].Value.([]byte)),
		SysUptime:  g.ToBigInt(result.Variables[1].Value),
		FreeMemory: g.ToBigInt(result.Variables[2].Value),
	}
}

func SnmpTrapListener(credentials resources.SnmpCredentials, s *server.Server) {
	params := &g.GoSNMP{
		Target:        credentials.Host,
		Port:          161,
		Version:       g.Version3,
		SecurityModel: g.UserSecurityModel,
		MsgFlags:      g.AuthPriv,
		Timeout:       time.Duration(30) * time.Second,
		SecurityParameters: &g.UsmSecurityParameters{UserName: credentials.User,
			AuthenticationProtocol:   g.MD5,
			AuthenticationPassphrase: credentials.Authentication,
			PrivacyProtocol:          g.AES,
			PrivacyPassphrase:        credentials.Privacy,
		},
	}

	listener := g.NewTrapListener()
	listener.Params = params
	listener.OnNewTrap = s.TrapHandler
	listener.Params.Logger = g.NewLogger(log.New(os.Stdout, "", 0))
	fmt.Println("Listening for traps...")
	listener.Listen(s.TrapHost)

}


//.1.3.6.1.4.1.9.9.41.2.0.1
//.1.3.6.1.4.1.9.9.43.2.0.1
