package account

import (
  "bytes"
  "fmt"
  "encoding/json"
  "net/http"

  api_account "metagit.org/blizzlike/cmangos-api/cmangos/realmd/account"
)

func CreateAccount(url, u, e, p, r, t string) (api_account.AccountError, error) {
  var ae api_account.AccountError
  var a api_account.AccountInfo

  a.Username = u
  a.Email = e
  a.Password = p
  a.Repeat = r

  ajson, err := json.Marshal(a)
  req, err := http.NewRequest("POST",
    fmt.Sprintf("%s/account", url), bytes.NewBuffer(ajson))
  if err != nil {
    return ae, err
  }

  req.Header.Set("Authorization", fmt.Sprintf("Token %s", t))
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    ae = api_account.AccountError{true, true, true, true}
    return ae, err
  }
  defer resp.Body.Close()

  if resp.StatusCode == http.StatusInternalServerError {
    ae = api_account.AccountError{true, true, true, true}
    return ae, fmt.Errorf("Cannot create account due to a internal server error")
  }

  _ = json.NewDecoder(resp.Body).Decode(&ae)
  if resp.StatusCode == http.StatusCreated {
    return ae, nil
  }

  return ae, fmt.Errorf("Cannot create account due to conflicts")
}

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
