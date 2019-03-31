package config

// "io/ioutil"
import (
  "encoding/json"
  "net/http"
)

type ApiConfig struct {
  NeedInvite bool `json:"needInvite,omitempty"`
  Realmd string `json:"realmd,omitempty"`
}

func FetchConfig(url string) (ApiConfig, error) {
  var cfg ApiConfig
  req, err := http.NewRequest("GET", url + "/config", nil)
  if err != nil {
    return cfg, err
  }

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    return cfg, err
  }
  defer resp.Body.Close()

  if resp.StatusCode != 200 {
    return cfg, err
  }

  //body, _ := ioutil.ReadAll(resp.Body)
  _ = json.NewDecoder(resp.Body).Decode(&cfg)
  return cfg, nil
}
