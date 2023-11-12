package main

import (
    "github.com/yborisovstc/desr/logr"
    "github.com/yborisovstc/desr/des"
    "net/http"
    "time"
)

var model = des.Des{ }

func init() {
    loadConfig()
    logr.LogI("Configuration: ", config)
}

func main() {
    logr.LogI("DESR starting ...")

    // handle static assets
    mux := http.NewServeMux()
    files := http.FileServer(http.Dir(config.StaticPath))
    mux.Handle("/static/", http.StripPrefix("/static/", files))

    mux.HandleFunc("/", handleIndex)

    mux.HandleFunc("/sysspec/post", handleSspecPost)

    // Create server and start listening
    server := http.Server{
	Addr:  config.Addr,
	Handler: mux,
	ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
	WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
	MaxHeaderBytes: 1 << 20,
    }
    server.ListenAndServe()
}

func handleIndex(writer http.ResponseWriter, request *http.Request) {
    logr.LogI("handleIndex ")
    logr.LogI("model Outp: ", model.Outp)
    generateHtml(writer, model, "layout", "navbar")
}
