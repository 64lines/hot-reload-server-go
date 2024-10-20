package main

import (
  "io/ioutil"
  "encoding/json"
)

type Config struct {
  Location string `json:"location"`
  Port int `json:"port"`
}

func LoadConfig(filePath string) (*Config, error) {
  data, err := ioutil.ReadFile(filePath)

  if err != nil {
    return nil, err
  }

  var config Config

  if err := json.Unmarshal(data, &config); err != nil {
    return nil, err
  }

  return &config, nil
}
