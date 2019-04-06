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

func AuthenticateByToken(url, token string) (bool, error) {
  req, err := http.NewRequest("GET", url + "/account/auth", nil)
  if err != nil {
    return false, err
  }

  req.Header.Set("Authorization", fmt.Sprintf("Token %s", token))
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    return false, err
  }
  defer resp.Body.Close()

  if resp.StatusCode != http.StatusOK {
    return false, fmt.Errorf("Not authenticated")
  }

  return true, nil
}
