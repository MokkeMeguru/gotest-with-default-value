package fixture

import "github.com/MokkeMeguru/gotest-with-default-value/internal/domain/model"

var Profile_Normal = model.Profile{
	FirstName: "Meguru",
	LastName:  "Mokke",
	FavoriteFoods: []model.Food{
		"curry", "ramen",
	},
	Skill: map[string]int{
		"clojure": 80,
		"golang":  1,
	},
	Post: &model.Post{
		Position: "server",
		Field:    1,
	},
}

var Profile_AnotherMan = model.Profile{
	FirstName: "Potter",
	LastName:  "Harry",
	FavoriteFoods: []model.Food{
		"curry", "ramen",
	},
	Skill: map[string]int{
		"clojure": 80,
		"golang":  1,
	},
	Post: &model.Post{
		Position: "server",
		Field:    1,
	},
}

var Profile_EmptyMan = model.Profile{
	FirstName: "",
	LastName:  "",
	FavoriteFoods: []model.Food{
		"curry", "ramen",
	},
	Skill: map[string]int{
		"clojure": 80,
		"golang":  1,
	},
	Post: &model.Post{
		Position: "server",
		Field:    1,
	},
}
