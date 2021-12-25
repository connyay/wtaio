package main

import (
	_ "embed"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geo"
	"github.com/paulmach/orb/geojson"
)

//go:embed map.geojson
var mapGeoJSON []byte

func main() {
	fc := geojson.NewFeatureCollection()
	json.Unmarshal(mapGeoJSON, &fc)
	log.SetOutput(os.Stdout)
	here := orb.Point{-106, 39}
	closest, distance := findClosestFeature(here, fc.Features)
	log.Printf("The closest feature is %#v with a distance of %v", closest.Properties.MustString("name", "[UNKNOWN]"), distance)

}

func findClosestFeature(reference orb.Point, features []*geojson.Feature) (closest *geojson.Feature, distance float64) {
	start := time.Now()
	defer log.Printf("It took %s to loop through %d features", time.Since(start), len(features))
	for i, feat := range features {
		d := geo.Distance(reference, feat.Point())
		if distance == 0 || d < distance {
			closest = features[i]
			distance = d
		}
	}
	return
}
