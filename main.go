package main

import (
  "github.com/gorilla/mux"
  "net/http"
  "go-htmx/src/pages"
)

func main() {
  gRouter := mux.NewRouter()
  gRouter.HandleFunc("/", pages.Homepage)
  gRouter.HandleFunc("/click", pages.ClickHandler).Methods("GET")
  http.ListenAndServe(":80", gRouter)
}