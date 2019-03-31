package config

import (
  ini "gopkg.in/ini.v1"

  apiconfig "metagit.org/blizzlike/cmangos-website/cmangos/api/config"
)

type Config struct {
  Host string
  Port int
  Api string
  Title string
  Templates string
  Static string
  NeedInvite bool
  Realmd string
}

var Cfg Config

func Read(file string) (Config, error) {
  c, err := ini.Load(file)
  if err != nil {
    return Cfg, err
  }

  Cfg.Host = c.Section("server").Key("listen").MustString("127.0.0.1")
  Cfg.Port = c.Section("server").Key("port").MustInt(5557)
  Cfg.Title = c.Section("server").Key("title").MustString("cmangos-website")

  Cfg.Templates = c.Section("paths").Key("templates").MustString("./templates")
  Cfg.Static = c.Section("paths").Key("public").MustString("./public")

  Cfg.Api = c.Section("api").Key("url").MustString("http://127.0.0.1:5556")

  var apicfg apiconfig.ApiConfig
  apicfg, err = apiconfig.FetchConfig(Cfg.Api)
  if err != nil {
    return Cfg, err
  }

  Cfg.NeedInvite = apicfg.NeedInvite
  Cfg.Realmd = apicfg.Realmd

  return Cfg, nil
}
