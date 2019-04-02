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

func RenderPost(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  u := strings.Join(r.Form["username"], "")
  p := strings.Join(r.Form["password"], "")

  token, err := auth.Authenticate(config.Cfg.Api, u, p)
  if err != nil {
    w.Header().Add("Location", "/login")
    w.WriteHeader(http.StatusFound)
    return
  }

  cookie := http.Cookie{
    Name: "auth",
    Value: token,
    Path: "/",
    MaxAge: config.Cfg.CookieMaxAge,
    HttpOnly: true,
  }
  http.SetCookie(w, &cookie)

  w.Header().Add("Location", "/dashboard")
  w.WriteHeader(http.StatusFound)
  return
}
