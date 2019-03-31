package landingpage

import (
  "html/template"
  "net/http"

  "metagit.org/blizzlike/cmangos-website/modules/config"
)

type LandingPageData struct {
  Title string
  LogonServer string
}

func Render(w http.ResponseWriter, r *http.Request) {
  tpl := template.Must(template.ParseFiles(
    config.Cfg.Templates + "/layout.html", config.Cfg.Templates + "/landingpage.html"))
  data := LandingPageData{
    Title: config.Cfg.Title,
    LogonServer: config.Cfg.Realmd,
  }
  w.WriteHeader(http.StatusOK)
  tpl.ExecuteTemplate(w, "layout", data)
  return
}
