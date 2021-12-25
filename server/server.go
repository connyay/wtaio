package server

import (
	"embed"
	"errors"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geo"
	"github.com/paulmach/orb/geojson"
)

var (
	//go:embed templates/*
	templates embed.FS
	//go:embed assets/*
	assets embed.FS
)

func ListenHTTP(addr string, features []*geojson.Feature) error {
	r := gin.Default()
	r.SetHTMLTemplate(template.Must(template.ParseFS(templates, "templates/*")))
	r.StaticFS("/static", http.FS(assets))
	r.GET("/:latitude/:longitude", func(c *gin.Context) {
		latitude, err := strconv.ParseFloat(c.Param("latitude"), 64)
		if err != nil {
			c.Error(err)
			return
		}
		longitude, err := strconv.ParseFloat(c.Param("longitude"), 64)
		if err != nil {
			c.Error(err)
			return
		}
		closest, _ := findClosestFeature(orb.Point{longitude, latitude}, features)
		if closest == nil {
			if err != nil {
				c.Error(errors.New("no feature close"))
				return
			}
		}
		c.Writer.WriteString(closest.Properties.MustString("name"))
	})
	r.GET("/", func(c *gin.Context) {
		data := map[string]interface{}{}
		c.HTML(http.StatusOK, "index.html.tmpl", data)
	})
	log.Printf("listening on http://%s", addr)
	return r.Run(addr)
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
