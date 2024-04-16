package dtos

type TeamsResponse struct {
	Message string      `json:"message"`
	Teams   interface{} `json:"teams"`
}

type Team struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
