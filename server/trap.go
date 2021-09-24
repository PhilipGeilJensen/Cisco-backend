package server

import (
	"Backend/resources"
	"fmt"
	g "github.com/gosnmp/gosnmp"
	"log"
	"net"
)

// Handles the incoming trap from the device - and makes sure it is send to the right recipients
func (s *Server) TrapHandler(packet *g.SnmpPacket, u *net.UDPAddr) {
	sliceOfOids := []string{".1.3.6.1.4.1.9.9.41.2.0.1"}
	fmt.Println(u.IP)
	t := TrapToStruct(packet)
	t.Host = u.IP.String()
	fmt.Println("Sending trap to chan")
	if sliceContains(sliceOfOids, t.Identifier) {
		for _, socketConn := range s.Websockets {
			if socketConn.Host == t.Host {
				socketConn.Conn.WriteJSON(t)
			}
		}
	}
}

// Simple function to check if a slice contains a value
func sliceContains(slice []string, obj string) bool {
	for _, o := range slice {
		if o == obj {
			return true
		}
	}
	return false
}

// Removes the connection from the slice
func RemoveConn(items []SocketClient, item SocketClient) []SocketClient {
	newitems := []SocketClient{}

	for _, i := range items {
		if i != item {
			newitems = append(newitems, i)
		}
	}

	return newitems
}

// Takes the trap and formats it to a struct
func TrapToStruct(s *g.SnmpPacket) resources.TrapObject {
	var t resources.TrapObject
	for i, variable := range s.Variables {
		var val string
		switch variable.Type {
		case g.OctetString:
			val = string(variable.Value.([]byte))
			t.Values = append(t.Values, val)
		case g.ObjectIdentifier:
			val = fmt.Sprintf("%s", variable.Value)
			t.Identifier = val
		case g.TimeTicks:
			a := g.ToBigInt(variable.Value)
			val = fmt.Sprintf("%d", (*a).Int64())
			t.Time = a
		case g.Null:
			val = ""
		default:
			// ... or often you're just interested in numeric values.
			// ToBigInt() will return the Value as a BigInt, for plugging
			// into your calculations.
			a := g.ToBigInt(variable.Value)
			val = fmt.Sprintf("%d", (*a).Int64())
		}
		log.Printf("- oid[%d]: %s (%s) = %v \n", i, variable.Name, variable.Type, val)
	}
	return t
}
