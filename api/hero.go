package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// Hero represents the attributes of a hero
type Hero struct {
	Name       string `json:"name"`
	Tag        string `json:"tag"`
	WeaponType string `json:"weapon_type"`
	MoveType   string `json:"move_type"`
	Rarity     []int  `json:"rarity"`

	BaseStats     []Stats     `json:"base"`
	BaseStatRange []StatRange `json:"base_stats,omitempty"`
	MaxStatRange  []StatRange `json:"max_stats,omitempty"`
	GrowthPoints  Stats       `json:"growth_points"`
}

// newHero reads in the json file for the hero name provided
// and returns a Hero object for that Hero
func newHero(name string) (*Hero, error) {
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

func (h *Hero) calcMaxStats() {
	for i, rarity := range h.Rarity {
		var base = StatRange{
			[]int{h.BaseStats[i].HP - 1, h.BaseStats[i].HP, h.BaseStats[i].HP + 1},
			[]int{h.BaseStats[i].Atk - 1, h.BaseStats[i].Atk, h.BaseStats[i].Atk + 1},
			[]int{h.BaseStats[i].Spd - 1, h.BaseStats[i].Spd, h.BaseStats[i].Spd + 1},
			[]int{h.BaseStats[i].Def - 1, h.BaseStats[i].Def, h.BaseStats[i].Def + 1},
			[]int{h.BaseStats[i].Res - 1, h.BaseStats[i].Res, h.BaseStats[i].Res + 1},
		}
		h.BaseStatRange = append(h.BaseStatRange, base)
		var max = StatRange{
			[]int{
				getGrowthValue(rarity, h.GrowthPoints.HP-1) + h.BaseStatRange[i].HP[0],
				getGrowthValue(rarity, h.GrowthPoints.HP) + h.BaseStatRange[i].HP[1],
				getGrowthValue(rarity, h.GrowthPoints.HP+1) + h.BaseStatRange[i].HP[2],
			},
			[]int{
				getGrowthValue(rarity, h.GrowthPoints.Atk-1) + h.BaseStatRange[i].Atk[0],
				getGrowthValue(rarity, h.GrowthPoints.Atk) + h.BaseStatRange[i].Atk[1],
				getGrowthValue(rarity, h.GrowthPoints.Atk+1) + h.BaseStatRange[i].Atk[2],
			},
			[]int{
				getGrowthValue(rarity, h.GrowthPoints.Spd-1) + h.BaseStatRange[i].Spd[0],
				getGrowthValue(rarity, h.GrowthPoints.Spd) + h.BaseStatRange[i].Spd[1],
				getGrowthValue(rarity, h.GrowthPoints.Spd+1) + h.BaseStatRange[i].Spd[2],
			},
			[]int{
				getGrowthValue(rarity, h.GrowthPoints.Def-1) + h.BaseStatRange[i].Def[0],
				getGrowthValue(rarity, h.GrowthPoints.Def) + h.BaseStatRange[i].Def[1],
				getGrowthValue(rarity, h.GrowthPoints.Def+1) + h.BaseStatRange[i].Def[2],
			},
			[]int{
				getGrowthValue(rarity, h.GrowthPoints.Res-1) + h.BaseStatRange[i].Res[0],
				getGrowthValue(rarity, h.GrowthPoints.Res) + h.BaseStatRange[i].Res[1],
				getGrowthValue(rarity, h.GrowthPoints.Res+1) + h.BaseStatRange[i].Res[2],
			},
		}
		h.MaxStatRange = append(h.MaxStatRange, max)
	}
}

func getHero(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	name := strings.ToLower(params["name"])
	h, err := newHero(name)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(h)
	}
}

func getMaxStats(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	name := strings.ToLower(params["name"])
	h, err := newHero(name)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		h.calcMaxStats()
		if r, ok := params["rarity"]; ok {
			rarity, _ := strconv.Atoi(r)
			for i, heroRarity := range h.Rarity {
				if heroRarity == rarity {
					json.NewEncoder(w).Encode(h.MaxStatRange[i])
					break
				}
				if i == len(h.Rarity)-1 {
					w.WriteHeader(http.StatusNotFound)
				}
			}
		} else {
			json.NewEncoder(w).Encode(h.MaxStatRange)
		}
	}
}
