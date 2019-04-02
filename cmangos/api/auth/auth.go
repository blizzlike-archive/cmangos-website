package auth

import (
  "fmt"
  "net/http"
)

func Authenticate(url, username, password string) (string, error) {
  req, err := http.NewRequest("POST", url + "/account/auth", nil)
  req.SetBasicAuth(username, password)
  if err != nil {
    return "", err
  }

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    return "", err
  }
  defer resp.Body.Close()

  if resp.StatusCode != http.StatusOK {
    return "", fmt.Errorf("Not authenticated")
  }

  token := resp.Header.Get("X-Auth-Token")
  return token, nil
}
