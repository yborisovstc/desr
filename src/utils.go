package main

import (
    "os"
    "log"
    "fmt"
    "encoding/json"
    "net/http"
    "html/template"
)

type Config struct {
    Addr string
    ReadTimeout int64
    WriteTimeout int64
    StaticPath string
}

var config Config

func loadConfig() {
    file, err := os.Open("config.json")
    if err != nil {
	log.Fatalln("Cannot open config file: ", err)
    }
    decoder := json.NewDecoder(file)
    config = Config{}
    err = decoder.Decode(&config)
    if err != nil {
	log.Fatalln("Failed parsing config file: ", err)
    }
}

func generateHtml(writer http.ResponseWriter, data interface{}, filenames ...string) {
    var fnames []string
    for _,fname := range filenames {
	fnames = append(fnames, fmt.Sprintf("templates/%s.html", fname))
    }
    templates := template.Must(template.ParseFiles(fnames...))
    templates.ExecuteTemplate(writer, "layout", data)
}




