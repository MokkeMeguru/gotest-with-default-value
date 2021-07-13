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
	Position string    `testdata:"beginner"`
	Field    PostField `testdata:"0"`
}

type Profile struct {
	FirstName     string         `testdata:"Meguru"`
	LastName      string         `testdata:"Mokke"`
	FavoriteFoods []Food         `testdata:"[\"curry\", \"ramen\"]"`
	Skill         map[string]int `testdata:"{\"clojure\": 80, \"golang\": 1}"`
	Post          Post
}

func (m Profile) Validate() error {
	if m.FirstName == "" || m.LastName == "" {
		return xerrors.New("you are ghost, aren't you?")
	}
	if m.FirstName != "Meguru" {
		return xerrors.New("who are you ?")
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
