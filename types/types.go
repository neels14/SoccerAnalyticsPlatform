package types

type SoccerMatchResponse struct {
	Year       int    `json:"year"`
	Stage      string `json:"stage"`
	HomeTeam   string `json:"home"`
	AwayTeam   string `json:"away"`
	Date       string `json:"date"`
	City       string `json:"city"`
	Stadium    string `json:"stadium"`
	Attendance int    `json:"attendance"`
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

type AverageGoalByCountry struct {
	Name        string  `json:"countryName"`
	AverageGoal float32 `json:"averageGoal"`
}

type TopScorer struct {
	Name      string `json:"countryName"`
	TopScorer string `json:"topScorer"`
	Goals     int    `json:"goals"`
}

type WorldCupAttendence struct {
	Year       int `json:"year"`
	Attendance int `json:"attendance"`
}

type Winner struct {
	Country string `json:"country_name"`
	Wins    int    `json:"total_win"`
}

type PlayerMostStarted struct {
	Name        string  `json:"playerName"`
	ShirtNumber int     `json:"shirtNumber"`
	Country     string  `json:"country"`
	StartRatio  float64 `json:"startRatio"`
}

type WinRatio struct {
	Country  string  `json:"country"`
	WinRatio float64 `json:"winRatio"`
}
