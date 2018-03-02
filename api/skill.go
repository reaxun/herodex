package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

// Skill ...
type Skill interface {
	getAvailability() int
}

// Weapon ...
type Weapon struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	Might     int    `json:"might"`
	Range     int    `json:"range"`
	Cost      int    `json:"cost"`
	Exclusive bool   `json:"exclusive"`
	Prereq    string `json:"prerequisite,omitempty"`
	Text      string `json:"text"`
}

func (w *Weapon) getAvailability() int {
	return 4
}

// Assist ...
type Assist struct {
	Name       string   `json:"name"`
	Range      int      `json:"range"`
	Cost       int      `json:"cost"`
	Exclusions []string `json:"exclusions"`
	Prereq     string   `json:"prerequisite,omitempty"`
	Text       string   `json:"text"`
}

func (a *Assist) getAvailability() int {
	return 4
}

// Special ...
type Special struct {
	Name       string   `json:"name"`
	Cooldown   int      `json:"cooldown"`
	Cost       int      `json:"cost"`
	Exclusions []string `json:"exclusions"`
	Exclusive  bool     `json:"exclusive"`
	Prereq     string   `json:"prerequisite,omitempty"`
	Text       string   `json:"text"`
}

func (s *Special) getAvailability() int {
	return 4
}

func newSkill(name string) (interface{}, error) {
	if _, err := os.Stat(fmt.Sprintf("assets/skills/weapons/%s.json", name)); err == nil {
		var w Weapon
		raw, _ := ioutil.ReadFile(fmt.Sprintf("assets/skills/weapons/%s.json", name))
		json.Unmarshal(raw, &w)
		return &w, nil
	} else if _, err := os.Stat(fmt.Sprintf("assets/skills/assists/%s.json", name)); err == nil {
		var a Assist
		raw, _ := ioutil.ReadFile(fmt.Sprintf("assets/skills/assists/%s.json", name))
		json.Unmarshal(raw, &a)
		return &a, nil
	} else if _, err := os.Stat(fmt.Sprintf("assets/skills/specials/%s.json", name)); err == nil {
		var s Special
		raw, _ := ioutil.ReadFile(fmt.Sprintf("assets/skills/specials/%s.json", name))
		json.Unmarshal(raw, &s)
		return &s, nil
	}
	return nil, fmt.Errorf("Skill %s not found", name)
}

func getSkill(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	name := strings.ToLower(params["name"])
	s, err := newSkill(name)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(s)
	}
}
