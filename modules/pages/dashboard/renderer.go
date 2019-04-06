package dashboard

import (
  "fmt"
  "os"
  "html/template"
  "net/http"

  api_character "metagit.org/blizzlike/cmangos-api/cmangos/mangosd/character"
  api_realm "metagit.org/blizzlike/cmangos-api/cmangos/realmd/realm"
  api_invite "metagit.org/blizzlike/cmangos-api/cmangos/api/account"

  "metagit.org/blizzlike/cmangos-website/modules/auth"
  "metagit.org/blizzlike/cmangos-website/modules/config"

  "metagit.org/blizzlike/cmangos-website/cmangos/api/character"
  "metagit.org/blizzlike/cmangos-website/cmangos/api/realm"
  "metagit.org/blizzlike/cmangos-website/cmangos/api/invite"
)

type Characterlist struct {
  Realm api_realm.Realm
  Characters []api_character.CharacterInfo
}

type DashboardPageData struct {
  Title string
  Discord string
  Realms []Characterlist
  Account int64
  InviteToken []api_invite.InviteInfo
}

func Render(w http.ResponseWriter, r *http.Request) {
  account, err := auth.Authenticate(w, r)
  if err != nil {
    return
  }

  var characterlist []Characterlist
  realms, err := realm.FetchRealms(config.Settings.Api)
  if err != nil {
    return
  }

  cookie, _ := r.Cookie("auth-token")
  token := cookie.Value

  tokens, _ := invite.GetInviteTokens(config.Settings.Api, token)

  for _, v := range realms {
    var cl Characterlist
    cl.Realm = v
    cl.Characters, err = character.FetchCharacters(
      config.Settings.Api, cl.Realm.Id, account.Id, token)
    fmt.Fprintf(os.Stdout, "Fetched %d characters for %s\n",
      len(cl.Characters), account.Username)
    if err != nil {
      continue
    }

    characterlist = append(characterlist, cl)
  }

  tpl, _ := template.ParseFiles(
    config.Settings.Templates + "/layout.html",
    config.Settings.Templates + "/header_small.html",
    config.Settings.Templates + "/dashboard.html")
  data := DashboardPageData{
    Title: config.Settings.Title,
    Discord: config.Settings.Discord,
    Realms: characterlist,
    Account: account.Id,
    InviteToken: tokens,
  }

  w.WriteHeader(http.StatusOK)
  tpl.Execute(w, data)
  return
}
