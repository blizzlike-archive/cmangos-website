package auth

import (
  "net/http"

  api_auth "metagit.org/blizzlike/cmangos-website/cmangos/api/auth"
  api_account "metagit.org/blizzlike/cmangos-api/cmangos/realmd/account"
  "metagit.org/blizzlike/cmangos-website/modules/config"
  a_account "metagit.org/blizzlike/cmangos-website/cmangos/api/account"
)

func Authenticated(r *http.Request) int64 {
  cookie, err := r.Cookie("auth-token")
  if err != nil {
    return 0
  }

  a, err := a_account.GetAccount(config.Settings.Api, cookie.Value)
  return a.Id
}

func Authenticate(w http.ResponseWriter, r *http.Request) (api_account.AccountInfo, error) {
  var account api_account.AccountInfo
  cookie, err := r.Cookie("auth-token")
  if err != nil {
    w.Header().Add("Location", "/login")
    w.WriteHeader(http.StatusFound)
    return account, err
  }

  verify, err := api_auth.AuthenticateByToken(config.Settings.Api, cookie.Value)
  if !verify {
    w.Header().Add("Location", "/login")
    w.WriteHeader(http.StatusFound)
    return account, err
  }

  account, err = a_account.GetAccount(config.Settings.Api, cookie.Value)
  if err != nil {
    return account, err
  }

  return account, nil
}
