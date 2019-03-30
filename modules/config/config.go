package config

import (
  ini "gopkg.in/ini.v1"
)

type Config struct {
  Host string
  Port int
  Api string
  Title string
  Templates string
  Static string
}

func Read(file string) (Config, error) {
  var cfg Config
  c, err := ini.Load(file)
  if err != nil {
    return cfg, err
  }

  cfg.Host = c.Section("server").Key("listen").MustString("127.0.0.1")
  cfg.Port = c.Section("server").Key("port").MustInt(5557)
  cfg.Title = c.Section("server").Key("title").MustString("cmangos-website")

  cfg.Templates = c.Section("paths").Key("templates").MustString("./templates")
  cfg.Static = c.Section("paths").Key("public").MustString("./public")

  cfg.Api = c.Section("api").Key("url").MustString("http://127.0.0.1:5556")

  return cfg, nil
}
