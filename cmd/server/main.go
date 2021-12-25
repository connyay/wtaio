package main

import (
	_ "embed"
	"encoding/json"
	"log"
	"os"

	"github.com/connyay/wtaio/server"
	"github.com/paulmach/orb/geojson"
)

//go:embed map.geojson
var mapGeoJSON []byte

func main() {
	httpAddr := os.Getenv("HTTP_ADDR")
	if httpAddr == "" {
		httpAddr = "0.0.0.0:8080"
	}

	fc := geojson.NewFeatureCollection()
	if err := json.Unmarshal(mapGeoJSON, &fc); err != nil {
		panic(err)
	}

	log.Fatal(server.ListenHTTP(httpAddr, fc.Features))
}
