package login

import (
  "html/template"
  "net/http"

  "metagit.org/blizzlike/cmangos-website/modules/config"
)

type LoginPageData struct {
  Title string
  Realmd string
  Discord string
}

func Render(w http.ResponseWriter, r *http.Request) {
  tpl, _ := template.ParseFiles(
    config.Cfg.Templates + "/layout.html",
    config.Cfg.Templates + "/header_login.html",
    config.Cfg.Templates + "/login.html")
  data := LoginPageData{
    Title: config.Cfg.Title,
    Realmd: config.Cfg.Realmd,
    Discord: config.Cfg.Discord,
  }

  w.WriteHeader(http.StatusOK)
  tpl.Execute(w, data)
  return
}
