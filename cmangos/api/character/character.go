package character

import (
  "fmt"
  "encoding/json"
  "net/http"

  api_character "metagit.org/blizzlike/cmangos-api/cmangos/mangosd/character"
)

func FetchCharacters(url string, realm int, account int64, token string) ([]api_character.CharacterInfo, error) {
  var cl []api_character.CharacterInfo
  req, err := http.NewRequest("GET",
    fmt.Sprintf("%s/realm/%d/characters/%d", url, realm, account), nil)
  if err != nil {
    return cl, err
  }

  req.Header.Set("Authorization", fmt.Sprintf("Token %s", token))
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    return cl, err
  }
  defer resp.Body.Close()

  if resp.StatusCode != 200 {
    return cl, err
  }

  _ = json.NewDecoder(resp.Body).Decode(&cl)
  return cl, nil
}
