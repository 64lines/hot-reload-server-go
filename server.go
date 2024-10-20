package main

import (
  "fmt"
  "io/ioutil"
  "encoding/json"
  "log"
  "net/http"
  "time"
)

type Response struct {
  Update bool `json:"update"`
  Message string `json:"message"`
}

func RunServer(config *Config) {
  currentTime := time.Now()

  http.HandleFunc("/hot-reload", func(w http.ResponseWriter, r*http.Request) {

    files, err := ioutil.ReadDir(config.Location)

    if err != nil {
      log.Fatal(err)
    }

    // Create a simple response object
    response := Response{
      Update: false,
      Message: "Nothing changed",
    }

    for _, file := range files {
      modificationTime := file.ModTime()

      if modificationTime.After(currentTime) {
        fmt.Println("File Modified: ", file.Name())
        currentTime = modificationTime

        response = Response{
          Update: true,
          Message: "File Modified: " + file.Name(),
        }
        break
      }
    }

    // Set the content type of the respose to JSON
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    if err := json.NewEncoder(w).Encode(response); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
  })

  log.Println("Starting server on :" + fmt.Sprintf("%d",config.Port))

  if err := http.ListenAndServe(":" + fmt.Sprintf("%d", config.Port), nil); err != nil {
    log.Fatal(err)
  }
}

