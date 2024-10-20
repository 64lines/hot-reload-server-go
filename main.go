package main

import (
  "log"
)

func main() {
  config, err := LoadConfig("config.json")

  if err != nil {
    log.Fatalf("Could not load config: %v", err)
  }
  RunServer(config)
}
