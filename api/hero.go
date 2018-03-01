package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// Stats act as a single struct for each of the 5 stats.
// This can be use for base stats, max stats, growth points, etc.
type Stats struct {
	HP  int `json:"hp"`
	Atk int `json:"atk"`
	Spd int `json:"spd"`
	Def int `json:"def"`
	Res int `json:"res"`
}

// Hero represents the attributes of a hero
type Hero struct {
	Name         string  `json:"name"`
	Tag          string  `json:"tag"`
	WeaponType   string  `json:"weapon_type"`
	MoveType     string  `json:"move_type"`
	Rarity       []int   `json:"rarity"`
	BaseStats    []Stats `json:"base_stats"`
	GrowthPoints Stats   `json:"growth_points"`
}

// NewHero reads in the json file for the hero name provided
// and returns a Hero object for that Hero
func NewHero(name string) (*Hero, error) {
	filename := fmt.Sprintf("assets/heroes/%s.json", name)
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	var h Hero
	json.Unmarshal(raw, &h)
	return &h, nil
}

func getHero(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	name := strings.ToLower(params["name"])
	h, err := NewHero(name)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(h)
	}
}
