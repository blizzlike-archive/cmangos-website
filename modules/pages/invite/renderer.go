package invite

import (
  "net/http"

  "metagit.org/blizzlike/cmangos-website/modules/auth"
  "metagit.org/blizzlike/cmangos-website/modules/config"

  "metagit.org/blizzlike/cmangos-website/cmangos/api/invite"
)

func RenderPost(w http.ResponseWriter, r *http.Request) {
  id := auth.Authenticated(r)
  if id == 0 {
    w.Header().Add("Location", "/login")
    w.WriteHeader(http.StatusFound)
    return
  }

  cookie, _ := r.Cookie("auth-token")
  token := cookie.Value

  invite.CreateInviteToken(config.Settings.Api, token)

  w.Header().Add("Location", "/dashboard")
  w.WriteHeader(http.StatusFound)
  return
}
