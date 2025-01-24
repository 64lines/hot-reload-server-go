/** Simple Hot reload server in Go */

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Location string `json:"location"`
	Port     int    `json:"port"`
	Filename string `json:"filename"`
}

const (
	HotReloadJSFile = "js/simple-hot-reload.js"
)

func main() {
	var config Config

	// Getting the configuration file.
	data, err := os.ReadFile("config.json")

	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	if err := json.Unmarshal(data, &config); err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	// Getting the JS code to inject in all files.
	code, err := os.ReadFile(HotReloadJSFile)

	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		data, err := os.ReadFile(fmt.Sprintf("%s/%s", config.Location, config.Filename))

		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(w, "%s<script>%s</script>", data, code)
	})

	log.Println("Starting server on :" + fmt.Sprintf("%d", config.Port))

	if err := http.ListenAndServe(":"+fmt.Sprintf("%d", config.Port), nil); err != nil {
		log.Fatal(err)
	}
}
