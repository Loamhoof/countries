package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"

	"github.com/paulmach/go.geo"
	"github.com/paulmach/go.geojson"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	f, err := os.Open("countries.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	decoder.DisallowUnknownFields()
	countries := make(Countries, 250)
	if err := decoder.Decode(&countries); err != nil {
		log.Fatal(err)
	}

	geos := make(map[string]*geojson.FeatureCollection)
	for _, country := range countries {
		geopath := filepath.Join("data", fmt.Sprintf("%s.geo.json", strings.ToLower(country.CCA3)))
		f, err := os.Open(geopath)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		decoder := json.NewDecoder(f)
		decoder.DisallowUnknownFields()
		geo := new(geojson.FeatureCollection)
		if err := decoder.Decode(geo); err != nil {
			log.Print(geopath)
			log.Fatal(err)
		}

		geos[country.CCA3] = geo
	}

	log.Print(polygonArea(geos["FRA"].Features[0].Geometry.MultiPolygon[9]) / 1e6)
}

func polygonArea(polygon [][][]float64) float64 {
	area := pathArea(polygon[0])

	if len(polygon) == 1 {
		return area
	}

	for _, hole := range polygon[1:] {
		area -= pathArea(hole)
	}

	return area
}

func pathArea(path [][]float64) float64 {
	var area float64
	for i := 0; i < len(path); i++ {
		area += rad(path[(i+1)%len(path)][0]-path[i][0]) * (2 + math.Sin(rad(path[i][1])) + math.Sin(rad(path[(i+1)%len(path)][1])))
	}

	return area * geo.EarthRadius * geo.EarthRadius / 2
}

func rad(x float64) float64 {
	return x * math.Pi / 180
}

type Countries []Country

type Country struct {
	AltSpellings []string          `json:"altSpellings"`
	Area         float64           `json:"area"`
	Borders      []string          `json:"borders"`
	CallingCode  []string          `json:"callingCode"`
	Capital      []string          `json:"capital"`
	CCA2         string            `json:"cca2"`
	CCA3         string            `json:"cca3"`
	CCN3         string            `json:"ccn3"`
	CIOC         string            `json:"cioc"`
	Currency     []string          `json:"currency"`
	Demonym      string            `json:"demonym"`
	Flag         string            `json:"flag"`
	Independent  bool              `json:"independent"`
	Landlocked   bool              `json:"landlocked"`
	Languages    map[string]string `json:"languages"`
	LatLng       [2]float64        `json:"latlng"`
	Name         struct {
		Common   string `json:"common"`
		Official string `json:"official"`
		Native   map[string]struct {
			Common   string `json:"common"`
			Official string `json:"official"`
		} `json:"native"`
	} `json:"name"`
	Region       string   `json:"region"`
	Status       string   `json:"status"`
	SubRegion    string   `json:"subregion"`
	TLD          []string `json:"tld"`
	Translations map[string]struct {
		Common   string `json:"common"`
		Official string `json:"official"`
	} `json:"translations"`
}
