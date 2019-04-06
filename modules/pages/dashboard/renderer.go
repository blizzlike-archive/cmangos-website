package dashboard

import (
  "html/template"
  "net/http"

  "metagit.org/blizzlike/cmangos-website/modules/config"
)

type DashboardPageData struct {
  Title string
  Discord string
}

func Render(w http.ResponseWriter, r *http.Request) {
  tpl, _ := template.ParseFiles(
    config.Settings.Templates + "/layout.html",
    config.Settings.Templates + "/header_small.html",
    config.Settings.Templates + "/dashboard.html")
  data := DashboardPageData{
    Title: config.Settings.Title,
    Discord: config.Settings.Discord,
  }

  w.WriteHeader(http.StatusOK)
  tpl.Execute(w, data)
  return
}
