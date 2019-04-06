package config

import (
  "encoding/json"
  "net/http"

  api_config "metagit.org/blizzlike/cmangos-api/cmangos/api/config"
)

func FetchConfig(url string) (api_config.ApiConfig, error) {
  var cfg api_config.ApiConfig
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

  _ = json.NewDecoder(resp.Body).Decode(&cfg)
  return cfg, nil
}
