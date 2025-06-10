package main

import (
  "github.com/gorilla/mux"
  "net/http"
  "text/template"
  "log"
)

func Homepage(w http.ResponseWriter, r *http.Request) {
  tmpl, _ := template.ParseFiles("templates/home.html")
  tmpl.ExecuteTemplate(w, "home.html", nil)
}

func main() {
  gRouter := mux.NewRouter()
  assetHandler := http.FileServer(http.Dir("./assets/"))
  gRouter.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", assetHandler))
  gRouter.HandleFunc("/", Homepage)
  log.Println("Server is running on http://localhost")
  http.ListenAndServe(":80", gRouter)
}