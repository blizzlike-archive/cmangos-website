package logout

import (
  "net/http"
  "time"
)

func Render(w http.ResponseWriter, r *http.Request) {
  cookie := &http.Cookie{
    Name: "auth-token",
    Value: "",
    Path: "/",
    Expires: time.Unix(0, 0),
    HttpOnly: true,
  }

  http.SetCookie(w, cookie)

  w.Header().Add("Location", "/")
  w.WriteHeader(http.StatusFound)
  return
}
