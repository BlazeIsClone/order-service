package user

import "strconv"

type User struct {
	ID     uint     `json:"id"`
	Name   string   `json:"name"`
	Powers []string `json:"powers"`
}

func users() []User {
	return []User{
		{
			ID:     1,
			Name:   "Dracula",
			Powers: []string{"Immortality", "Shape-shifting", "Mind Control"},
		},
		{
			ID:     2,
			Name:   "Frankenstein",
			Powers: []string{"Superhuman Strength", "Endurance"},
		},
		{
			ID:     3,
			Name:   "Werewolf",
			Powers: []string{"Shape-shifting", "Enhanced Senses", "Regeneration"},
		},
		{
			ID:     4,
			Name:   "Zombie",
			Powers: []string{"Undead Physiology", "Immunity to Pain"},
		},
		{
			ID:     5,
			Name:   "Mummy",
			Powers: []string{"Immortality", "Control over Sand"},
		},
	}
}

func loadUsers() map[string]User {
	users := users()
	res := make(map[string]User, len(users))

	for _, x := range users {
		res[strconv.Itoa(int(x.ID))] = x
	}

	return res
}
