package server

import (
	"net/http"

	gmux "github.com/gorilla/mux"
	ws "github.com/gorilla/websocket"
)

// The server object it self
// Most importantly it holds the list of connected websockets.
type Server struct {
	TrapHost   string
	Websockets []SocketClient
	mux        *gmux.Router
}

// A simple connection struct, for easy access to connection and the requested host
type SocketClient struct {
	Conn *ws.Conn
	Host string
}

// Create's a new instance of the server struct
func New(trapHost string) *Server {
	mux := gmux.NewRouter()
	var websockets []SocketClient
	return &Server{
		TrapHost:   trapHost,
		Websockets: websockets,
		mux:        mux,
	}
}

// Implentation of ServeHTTP - also handles CORS
func (srv *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	srv.mux.ServeHTTP(w, r)
}

// Simple metrics handler which holds the path and the handler to it
type metricsHandler struct {
	Path    string
	Handler http.Handler
}

// MetricsHandler implentation of ServeHTTP
func (m *metricsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.Handler.ServeHTTP(w, r)
}
