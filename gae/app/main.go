package main

import (
	"net/http"

	"github.com/kik4/my-portfolio-nuxt/gae/handler"
	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/qiita/items/request", handler.RequestQiitaItems)
	http.HandleFunc("/qiita/items", handler.ResponseQiitaItems)
	http.HandleFunc("/", handler.Hello)
	appengine.Main()
}
