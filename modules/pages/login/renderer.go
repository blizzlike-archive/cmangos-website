package login

import (
  "html/template"
  "net/http"
  "strings"

  "metagit.org/blizzlike/cmangos-website/cmangos/api/auth"

  "metagit.org/blizzlike/cmangos-website/modules/config"
  a_auth "metagit.org/blizzlike/cmangos-website/modules/auth"
)

type LoginPageData struct {
  Title string
  Realmd string
  Discord string
  Account int64
}

func RenderGet(w http.ResponseWriter, r *http.Request) {
  id := a_auth.Authenticated(r)

  tpl, _ := template.ParseFiles(
    config.Settings.Templates + "/layout.html",
    config.Settings.Templates + "/header_login.html",
    config.Settings.Templates + "/login.html")
  data := LoginPageData{
    Title: config.Settings.Title,
    Realmd: config.Settings.RealmdAddress,
    Discord: config.Settings.Discord,
    Account: id,
  }

  w.WriteHeader(http.StatusOK)
  tpl.Execute(w, data)
  return
}

func RenderPost(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  u := strings.Join(r.Form["username"], "")
  p := strings.Join(r.Form["password"], "")

  token, err := auth.Authenticate(config.Settings.Api, u, p)
  if err != nil {
    w.Header().Add("Location", "/login")
    w.WriteHeader(http.StatusFound)
    return
  }

  cookie := http.Cookie{
    Name: "auth-token",
    Value: token,
    Path: "/",
    MaxAge: config.Settings.CookieMaxAge,
    HttpOnly: true,
  }
  http.SetCookie(w, &cookie)

  w.Header().Add("Location", "/dashboard")
  w.WriteHeader(http.StatusFound)
  return
}
