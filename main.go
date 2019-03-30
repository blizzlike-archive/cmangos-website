package main

import (
  "fmt"
  "log"
  "os"
  "net/http"

  "github.com/gorilla/mux"

  "metagit.org/blizzlike/cmangos-website/modules/config"
  "metagit.org/blizzlike/cmangos-website/modules/pages/landingpage"
)

func main() {
  if len(os.Args) != 2 {
    fmt.Fprintf(os.Stderr, "USAGE: %s <config>\n", os.Args[0])
    os.Exit(1)
  }
  cfg, err := config.Read(os.Args[1])
  if err != nil {
    fmt.Fprintf(os.Stderr, "Failed to read file %v\n", err)
    os.Exit(2)
  }

  landingpage.Cfg = cfg

  router := mux.NewRouter()
  router.HandleFunc("/", landingpage.Render).Methods("GET")
  router.PathPrefix("/").Handler(
    http.StripPrefix("/static/", http.FileServer(http.Dir(cfg.Static))))

  log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port), router))
}