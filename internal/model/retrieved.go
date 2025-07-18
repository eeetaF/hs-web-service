package model

type RetrievedLeaderboardData struct {
	Leaderboard struct {
		Rows []struct {
			Rank      int    `json:"rank"`
			AccountID string `json:"accountid"`
			Rating    int    `json:"rating"`
		} `json:"rows"`
	} `json:"leaderboard"`
}
