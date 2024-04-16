package dtos

type Matches struct {
	FkPhase            int `json:"fkPhase"`
	FkLocalTeam        int `json:"fkLocalTeam"`
	FkVisitorTeam      int `json:"fkVisitorTeam"`
	GoalsLocal         int `json:"goalsLocal"`
	GoalsVisitor       int `json:"goalsVisitor"`
	YellowCardsLocal   int `json:"yellowCardsLocal"`
	YellowCardsVisitor int `json:"yellowCardsVisitor"`
	RedCardsLocal      int `json:"redCardsLocal"`
	RedCardsVisitor    int `json:"redCardsVisitor"`
	Winner             int `json:"winner"`
}
