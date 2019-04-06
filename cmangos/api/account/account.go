package account

import (
  "fmt"
  "encoding/json"
  "net/http"

  api_account "metagit.org/blizzlike/cmangos-api/cmangos/realmd/account"
)

func GetAccount(url, token string) (api_account.AccountInfo, error) {
  var a api_account.AccountInfo
  req, err := http.NewRequest("GET",
    fmt.Sprintf("%s/account", url), nil)
  if err != nil {
    return a, err
  }

  req.Header.Set("Authorization", fmt.Sprintf("Token %s", token))
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    return a, err
  }
  defer resp.Body.Close()

  if resp.StatusCode != 200 {
    return a, err
  }

  _ = json.NewDecoder(resp.Body).Decode(&a)
  return a, nil
}
