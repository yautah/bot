package data

type Result struct {
	WarType        string `json:"type"`
	UtcTime        int64  `json:utcTime`
	Winner         int64  `json:winner`
	TeamCrowns     int64  `json:teamCrowns`
	OpponentCrowns int64  `json:opponentCrowns`
	Team           []Team `json:team`
	Opponent       []Team `json:opponent`
}
