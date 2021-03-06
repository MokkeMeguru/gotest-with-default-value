package model

import (
	"fmt"

	"golang.org/x/xerrors"
)

type Food string

const (
	Curry  Food = "curry"
	Ramen  Food = "ramen"
	Buffet Food = "buffet"
)

type PostField int

const (
	ServerEngineer PostField = iota
	ClientEngineer
	Designer
	Planner
)

type Post struct {
	Position string    `json:"position" testdata:"beginner"`
	Field    PostField `json:"field" testdata:"0"`
}

type Profile struct {
	FirstName     string         `json:"first_name" testdata:"Meguru"`
	LastName      string         `json:"last_name" testdata:"Mokke"`
	FavoriteFoods []Food         `json:"favorite_foods" testdata:"[\"curry\", \"ramen\"]"`
	Skill         map[string]int `json:"skill" testdata:"{\"clojure\": 80, \"golang\": 1}"`
	Post          *Post          `json:"post" testdata:"{}"`
}

func (m Profile) Validate() error {
	if m.FirstName == "" || m.LastName == "" {
		return xerrors.New("you are ghost, aren't you?")
	}
	if m.FirstName != "Meguru" {
		return xerrors.New("who are you ?")
	}
	if m.Post == nil {
		return xerrors.New("no post?")
	}
	for language, level := range m.Skill {
		switch language {
		case "clojure", "javascript", "typescript", "golang":
		default:
			return xerrors.New(fmt.Sprintf("invalid language %v", language))
		}
		if 0 > level && level > 100 {
			return xerrors.New(fmt.Sprintf("invalid skill level %v", level))
		}
	}
	return nil
}

func (m Profile) FullName() (string, error) {
	return fmt.Sprintf("%s %s", m.LastName, m.FirstName), nil
}
