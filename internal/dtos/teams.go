package dtos

type TeamsOrdered struct {
	Position int    `json:"position"`
	Name     string `json:"name"`
	Points   int    `json:"points"`
}
