package landingpage

import (
  "html/template"
  "net/http"

  "metagit.org/blizzlike/cmangos-website/modules/config"
)

type LandingPageData struct {
  Title string
}

var Cfg config.Config

func Render(w http.ResponseWriter, r *http.Request) {
  tpl := template.Must(template.ParseFiles(
    Cfg.Templates + "/layout.html", Cfg.Templates + "/landingpage.html"))
  data := LandingPageData{
    Title: Cfg.Title,
  }
  w.WriteHeader(http.StatusOK)
  tpl.ExecuteTemplate(w, "layout", data)
  return
}
