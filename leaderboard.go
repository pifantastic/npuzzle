package main

type LeaderboardEntry struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Moves int    `json:"moves"`
	Hash  string `json:"hash"`
}

type Leaderboard struct {
	Entries []LeaderboardEntry `json:"entries"`
}
