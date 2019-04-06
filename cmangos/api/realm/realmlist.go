package realm

import (
  "encoding/json"
  "net/http"

  api_realm "metagit.org/blizzlike/cmangos-api/cmangos/realmd/realm"
)

func FetchRealms(url string) ([]api_realm.Realm, error) {
  var rl []api_realm.Realm
  req, err := http.NewRequest("GET", url + "/realm", nil)
  if err != nil {
    return rl, err
  }

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    return rl, err
  }
  defer resp.Body.Close()

  if resp.StatusCode != 200 {
    return rl, err
  }

  _ = json.NewDecoder(resp.Body).Decode(&rl)
  return rl, nil
}
