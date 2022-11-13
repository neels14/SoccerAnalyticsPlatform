package types

import "time"

type SoccerMatchResponse struct {
	StatusCode int       `json:"code"`
	Year       int       `json:"year"`
	Stage      string    `json:"stage"`
	HomeTeam   string    `json:"home"`
	AwayTeam   string    `json:"away"`
	Date       time.Time `json:"time"`
	City       string    `json:"city"`
	Stadium    string    `json:"stadium"`
	Attendance int       `json:"attendance"`
	HomeCoach  string    `json:"homeCoach"`
	AwayCoach  string    `json:"awayCoach"`
}

type CountryResponse struct {
	Name        string `json:"name"`
	TeamInitial string `json:"teamInitial"`
}

type PlayerResponse struct {
	Name        string `json:"name"`
	ShirtNumber int    `json:"shirtNumber"`
	Country     string `json:"country"`
}

type WorldCupResponse struct {
	Year        string `json:"year"`
	HostCountry string `json:"host"`
	FirstPlace  string `json:"first"`
	SecondPlace string `json:"second"`
	ThirdPlace  string `json:"third"`
}

type PlayerInMatch struct {
	PlayerName  string `json:"playerName"`
	ShirtNumber int    `json:"shirtNumber"`
	Country     string `json:"country"`
	Year        int    `json:"year"`
	Stage       string `json:"stage"`
	HomeTeam    string `json:"home"`
	AwayTeam    string `json:"away"`
	Starter     bool   `json:"isStarter"`
	GoalsScored int    `json:"goalsScored"`
}