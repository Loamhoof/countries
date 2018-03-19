package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	fCountries, err := os.Open("countries.json")
	if err != nil {
		log.Fatal(err)
	}
	defer fCountries.Close()

	decoder := json.NewDecoder(fCountries)
	decoder.DisallowUnknownFields()
	countries := make(Countries, 250)
	if err := decoder.Decode(&countries); err != nil {
		log.Fatal(err)
	}

	fDeck, err := os.Create("deck.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer fDeck.Close()

	w := csv.NewWriter(fDeck)
	defer w.Flush()

	w.Write(COLUMNS)

	log.Print(countries[0])
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

const COLUMNS = []string{
	"name.common",
	"name.official",
	"tlds",
	"cca2",
	"ccn3",
	"cca3",
	"cioc",
	"independent",
	"status",
	"currencies",
	"callingCodes",
	"capitals",
	"altSpellings",
	"region",
	"subregion",
	"languages",
	"translation.cym.common",
	"translation.cym.official",
	"translation.deu.common",
	"translation.deu.official",
	"translation.est.common",
	"translation.est.official",
	"translation.fin.common",
	"translation.fin.official",
	"translation.fra.common",
	"translation.fra.official",
	"translation.hrv.common",
	"translation.hrv.official",
	"translation.ita.common",
	"translation.ita.official",
	"translation.jpn.common",
	"translation.jpn.official",
	"translation.nld.common",
	"translation.nld.official",
	"translation.por.common",
	"translation.por.official",
	"translation.rus.common",
	"translation.rus.official",
	"translation.slk.common",
	"translation.slk.official",
	"translation.spa.common",
	"translation.spa.official",
	"translation.zho.common",
	"translation.zho.official",
	"demonym",
	"landlocked",
	"borders",
	"area",
	"flag",
}

func (c *Country) toRecord() []string {
	record := []string{
		c.Name.Common,
		c.Name.Official,
		ssts(c.TLD),
		c.CCA2,
		c.CCN3,
		c.CCA3,
		c.CIOC,
		bts(c.Independent),
		c.Status,
		ssts(c.Currency),
		ssts(c.CallingCode),
		ssts(c.Capital),
		ssts(c.AltSpellings),
		c.Region,
		c.SubRegion,
		mssts(c.Languages),
	}

	return record
}

func bts(b bool) string {
	if b {
		return "X"
	}

	return ""
}

func ssts(ss []string) string {
	return strings.Join(ss, "/")
}

func mssts(mss map[string]string) string {
	return ""
	// values := make([]string, len(mss))

	// i := 0
	// for _, v := range values {

	// }
}
