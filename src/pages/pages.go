package pages

import (
  "log"
  "net/http"
  "text/template"
  "go-htmx/src/components/button"
)

func Homepage(w http.ResponseWriter, r *http.Request) {
  tmpl, _ := template.ParseFiles("templates/home.html")
  components := map[string]string{
    "button": button.Button("Click Me", "/click"),
  }
  tmpl.ExecuteTemplate(w, "home.html", components)
}

func ClickHandler(w http.ResponseWriter, r *http.Request) {
  log.Println("Button clicked")
  tmpl, _ := template.ParseFiles("templates/content.html")
  tmpl.ExecuteTemplate(w, "content.html", nil)
}