package main

import (
	"Backend/api"
	"Backend/server"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"regexp"
	"time"
)

type httpHandler struct {
	Server *server.Server
}

// Pass ServeHttp to server instance
func (h *httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Server.ServeHTTP(w, r)
}

// Http handler where API i initiated and server is passed to it
func makeHTTPHandler(srv *server.Server) *httpHandler {
	api.Init(srv)

	var handler httpHandler
	handler.Server = srv
	return &handler
}

// Make a http server which listens on port 8080
func makeAndStartHTTPServer(srv *server.Server) {
	fmt.Println("Starting HTTP server")
	httpSrv = &http.Server{
		Addr:         ":8080",
		Handler:      makeHTTPHandler(srv),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	err := httpSrv.ListenAndServe()
	if err == http.ErrServerClosed {
		return
	}
	if err != nil {
		logrus.Fatalf("HTTP server error: %s", err)
	}
}

func makeAndServeFileServer() {
	fileServer := http.FileServer(http.Dir("public"))
	fileMatcher := regexp.MustCompile(`\.[a-zA-Z]*$`)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if !fileMatcher.MatchString(r.URL.Path) {
			http.ServeFile(w, r, "public/index.html")
		} else {
			fileServer.ServeHTTP(w, r)
		}
	})
	http.ListenAndServe(":80", nil)
}
