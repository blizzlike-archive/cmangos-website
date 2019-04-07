package signup

import (
  "html/template"
  "net/http"
  "strings"
  "fmt"
  "os"

  "metagit.org/blizzlike/cmangos-website/cmangos/api/auth"
  "metagit.org/blizzlike/cmangos-website/cmangos/api/account"

  "metagit.org/blizzlike/cmangos-website/modules/config"
  a_auth "metagit.org/blizzlike/cmangos-website/modules/auth"
  api_account "metagit.org/blizzlike/cmangos-api/cmangos/realmd/account"
)

type SignupPageData struct {
  Title string
  Realmd string
  Discord string
  Account int64
  AccountError api_account.AccountError
}

func RenderPage(w http.ResponseWriter, r *http.Request, ae api_account.AccountError) {
  id := a_auth.Authenticated(r)
  tpl, _ := template.ParseFiles(
    config.Settings.Templates + "/layout.html",
    config.Settings.Templates + "/header_signup.html",
    config.Settings.Templates + "/login.html")
  data := SignupPageData{
    Title: config.Settings.Title,
    Realmd: config.Settings.RealmdAddress,
    Discord: config.Settings.Discord,
    Account: id,
    AccountError: ae,
  }

  w.WriteHeader(http.StatusOK)
  tpl.Execute(w, data)
}

func RenderGet(w http.ResponseWriter, r *http.Request) {
  ae := api_account.AccountError{true, true, true, true}
  RenderPage(w, r, ae)
  return
}

func RenderPost(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  u := strings.Join(r.Form["username"], "")
  e := strings.Join(r.Form["email"], "")
  p := strings.Join(r.Form["password"], "")
  rp := strings.Join(r.Form["repeat"], "")
  t := strings.Join(r.Form["token"], "")

  ae, err := account.CreateAccount(
    config.Settings.Api, u, e, p, rp, t)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Cannot create account (%v)\n", err)
    RenderPage(w, r, ae)
    return
  }

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
