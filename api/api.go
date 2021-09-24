package api

import (
	"Backend/server"
)

type API struct {
	srv *server.Server
}

func Init(srv *server.Server) {
	api := API{
		srv: srv,
	}
	// GET Methods
	srv.API("GET", "/api/hostname/{hostname}", ChangeHostname)
	srv.API("GET", "/api/ws/{host}", api.Websocket)

	// POST Methods
	srv.API("POST", "/api/save", SaveConfig)
	srv.API("POST", "/api/snmp", GetSnmpInfo)
	srv.API("POST", "/api/connect", CreateConnection)
	srv.API("POST", "/api/interface/configure", ConfigureInterface)
}

