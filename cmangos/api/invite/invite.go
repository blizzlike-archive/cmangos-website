package invite

import (
  "fmt"
  "encoding/json"
  "net/http"

  api_invite "metagit.org/blizzlike/cmangos-api/cmangos/api/account"
)

func CreateInviteToken(url, token string) {
  req, err := http.NewRequest("POST",
    fmt.Sprintf("%s/account/invite", url), nil)
  if err != nil {
    return
  }

  req.Header.Set("Authorization", fmt.Sprintf("Token %s", token))
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    return
  }
  defer resp.Body.Close()

  if resp.StatusCode != http.StatusCreated {
    return
  }

  return
}

func GetInviteTokens(url, token string) ([]api_invite.InviteInfo, error) {
  var ii []api_invite.InviteInfo
  req, err := http.NewRequest("GET",
    fmt.Sprintf("%s/account/invite", url), nil)
  if err != nil {
    return ii, err
  }

  req.Header.Set("Authorization", fmt.Sprintf("Token %s", token))
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    return ii, err
  }
  defer resp.Body.Close()

  if resp.StatusCode != 200 {
    return ii, err
  }

  _ = json.NewDecoder(resp.Body).Decode(&ii)
  return ii, nil
}
