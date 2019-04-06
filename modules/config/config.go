package config

import (
  ini "gopkg.in/ini.v1"

  api_config "metagit.org/blizzlike/cmangos-api/cmangos/api/config"
  a_config "metagit.org/blizzlike/cmangos-website/cmangos/api/config"
)

type Config struct {
  Host string
  Port int
  Api string
  Title string
  Templates string
  Static string
  NeedInvite bool
  RealmdAddress string
  RealmdPort int
  Discord string
  CookieMaxAge int
}

var Settings Config

func Read(file string) (Config, error) {
  c, err := ini.Load(file)
  if err != nil {
    return Settings, err
  }

  Settings.Host = c.Section("server").Key("listen").MustString("127.0.0.1")
  Settings.Port = c.Section("server").Key("port").MustInt(5557)
  Settings.Title = c.Section("server").Key("title").MustString("cmangos-website")
  Settings.Discord = c.Section("server").Key("discord").MustString("")
  Settings.CookieMaxAge = c.Section("server").Key("cookie").MustInt(60 * 60)

  Settings.Templates = c.Section("paths").Key("templates").MustString("./templates")
  Settings.Static = c.Section("paths").Key("public").MustString("./public")

  Settings.Api = c.Section("api").Key("url").MustString("http://127.0.0.1:5556")

  var apicfg api_config.ApiConfig
  apicfg, err = a_config.FetchConfig(Settings.Api)
  if err != nil {
    return Settings, err
  }

  Settings.NeedInvite = apicfg.RequireInvite
  Settings.RealmdAddress = apicfg.RealmdAddress
  Settings.RealmdPort = apicfg.RealmdPort

  return Settings, nil
}
