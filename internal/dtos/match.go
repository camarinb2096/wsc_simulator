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

type MatchDetail struct {
	Phase              int    `json:"phase"`
	TeamLocalName      string `json:"teamLocal"`
	TeamVisitorName    string `json:"teamVisitor"`
	GoalsLocal         int    `json:"goalsLocal"`
	GoalsVisitor       int    `json:"goalsVisitor"`
	YellowCardsLocal   int    `json:"yellowCardsLocal"`
	YellowCardsVisitor int    `json:"yellowCardsVisitor"`
	RedCardsLocal      int    `json:"redCardsLocal"`
	RedCardsVisitor    int    `json:"redCardsVisitor"`
	Winner             string `json:"winner"`
}

type MatchStatistics struct {
	Name          string `json:"name"`
	GoalsScored   int    `json:"goalsScored"`
	GoalsConceded int    `json:"goalsConceded"`
}
type MatchResponse struct {
	Message string      `json:"message"`
	Total   int         `json:"total"`
	Data    interface{} `json:"data"`
}
