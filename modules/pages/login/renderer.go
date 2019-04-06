package login

import (
  "html/template"
  "net/http"
  "strings"

  "metagit.org/blizzlike/cmangos-website/cmangos/api/auth"

  "metagit.org/blizzlike/cmangos-website/modules/config"
)

type LoginPageData struct {
  Title string
  Realmd string
  Discord string
}

func RenderGet(w http.ResponseWriter, r *http.Request) {
  tpl, _ := template.ParseFiles(
    config.Settings.Templates + "/layout.html",
    config.Settings.Templates + "/header_login.html",
    config.Settings.Templates + "/login.html")
  data := LoginPageData{
    Title: config.Settings.Title,
    Realmd: config.Settings.RealmdAddress,
    Discord: config.Settings.Discord,
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
