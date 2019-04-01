package realm

import (
  "encoding/json"
  "net/http"

  "metagit.org/blizzlike/cmangos-api/cmangos/iface"
)

func FetchRealms(url string) (iface.Realmlist, error) {
  var rl iface.Realmlist
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
