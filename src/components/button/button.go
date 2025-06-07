package button

func Button(text string, href string) string {
  return `<button hx-get="` + href + `" hx-target="#content" class="btn btn-primary">` + text + `</button>`
}