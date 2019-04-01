package landingpage

import (
  "fmt"
  "html/template"
  "net/http"
  "os"

  "metagit.org/blizzlike/cmangos-api/cmangos/iface"

  "metagit.org/blizzlike/cmangos-website/modules/config"
  "metagit.org/blizzlike/cmangos-website/cmangos/api/realm"
)

type LandingPageData struct {
  Title string
  Realmd string
  Realmlist []iface.Realm
  Discord string
}

func Render(w http.ResponseWriter, r *http.Request) {
  rl, err := realm.FetchRealms(config.Cfg.Api)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Cannot fetch realmlist (%v)\n", err)
  }

  tpl := template.Must(template.ParseFiles(
    config.Cfg.Templates + "/layout.html", config.Cfg.Templates + "/landingpage.html"))
  data := LandingPageData{
    Title: config.Cfg.Title,
    Realmd: config.Cfg.Realmd,
    Realmlist: rl.Realmlist,
    Discord: config.Cfg.Discord,
  }

  w.WriteHeader(http.StatusOK)
  tpl.ExecuteTemplate(w, "layout", data)
  return
}
