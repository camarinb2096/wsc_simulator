package dtos

type Groups struct {
	GroupA []Group `json:"groupA"`
	GroupB []Group `json:"groupB"`
	GroupC []Group `json:"groupC"`
	GroupD []Group `json:"groupD"`
}

type Group struct {
	FkTeam1 int `json:"fkTeam1"`
	FkTeam2 int `json:"fkTeam2"`
	FkTeam3 int `json:"fkTeam3"`
	FkTeam4 int `json:"fkTeam4"`
}
