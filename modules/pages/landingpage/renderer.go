package landingpage

import (
  "fmt"
  "html/template"
  "net/http"
  "os"

  api_realm "metagit.org/blizzlike/cmangos-api/cmangos/realmd/realm"

  "metagit.org/blizzlike/cmangos-website/modules/config"
  "metagit.org/blizzlike/cmangos-website/modules/auth"
  "metagit.org/blizzlike/cmangos-website/cmangos/api/realm"
)

type LandingPageData struct {
  Title string
  Realmd string
  Realmlist []api_realm.Realm
  Discord string
  Account int64
}

func Render(w http.ResponseWriter, r *http.Request) {
  id := auth.Authenticated(r)

  rl, err := realm.FetchRealms(config.Settings.Api)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Cannot fetch realmlist (%v)\n", err)
  }

  tpl, _ := template.ParseFiles(
    config.Settings.Templates + "/layout.html",
    config.Settings.Templates + "/landingpage.html",
    config.Settings.Templates + "/header_landingpage.html")
  data := LandingPageData{
    Title: config.Settings.Title,
    Realmd: config.Settings.RealmdAddress,
    Realmlist: rl,
    Discord: config.Settings.Discord,
    Account: id,
  }

  w.WriteHeader(http.StatusOK)
  tpl.Execute(w, data)
  return
}
