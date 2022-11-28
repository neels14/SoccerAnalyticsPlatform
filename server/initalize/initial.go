package initalize

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/qustavo/dotsql"
)

var lockDOT = &sync.Mutex{}
var lockDB = &sync.Mutex{}
var initalizedDOT *dotsql.DotSql
var initalizedDB *sql.DB

const testDBName string = "testdb"
const prodDBName string = "proddb"

func csvMatches(wg *sync.WaitGroup, db *sql.DB, dot *dotsql.DotSql) {
	println("Started CSV MATCHES")
	defer wg.Done()
	dot.Exec(db, "CSV-DROP-WorldCupMatchesCSV")
	dot.Exec(db, "CSV-Table-WorldCupMatches")
	dot.Exec(db, "CSV-Table-WorldCupMatches-Insert")
	println("Done CSV MATCHES")
}

func csvPlayers(wg *sync.WaitGroup, db *sql.DB, dot *dotsql.DotSql) {
	println("Started CSV PLAYERS")
	defer wg.Done()
	dot.Exec(db, "CSV-DROP-WorldCupPlayersCSV")
	dot.Exec(db, "CSV-Table-WorldCupPlayersCSV")
	dot.Exec(db, "CSV-Table-WorldCupPlayers-Insert")
	println("Done CSV PLAYERS")
}

func csvWorldCups(wg *sync.WaitGroup, db *sql.DB, dot *dotsql.DotSql) {
	println("Started CSV WC")
	defer wg.Done()
	dot.Exec(db, "CSV-DROP-WorldCupsCSV")
	dot.Exec(db, "CSV-Table-WorldCupsCSV")
	dot.Exec(db, "CSV-Table-WorldCupsCSV-Insert")
	println("Done CSV WC")
}

func GetDB() *sql.DB {
	if initalizedDB == nil {
		lockDB.Lock()
		defer lockDB.Unlock()
		if initalizedDB == nil {
			fmt.Println("Creating DB now.")
			db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/"+testDBName)
			if err != nil {
				panic(err.Error())
			}
			initalizedDB = db
		} else {
			fmt.Println("DB made already")
		}

	}
	return initalizedDB
}

func GetInstance() *dotsql.DotSql {
	if initalizedDOT == nil {
		lockDOT.Lock()
		defer lockDOT.Unlock()
		if initalizedDOT == nil {
			fmt.Println("Creating DOT now.")
			db := GetDB()
			//Loads queries from file
			dotQueryCountry, errQueryCountry := dotsql.LoadFromFile("./sql/queries/generalQueries.sql")
			if errQueryCountry != nil {
				panic(errQueryCountry.Error())
			}
			dotDDLCountries, errDDLCountries := dotsql.LoadFromFile("./sql/ddl/country.sql")
			if errDDLCountries != nil {
				panic(errDDLCountries.Error())
			}
			dotDDLPlayersInMatch, errDDLPlayersInMatch := dotsql.LoadFromFile("./sql/ddl/player_plays_in_match.sql")
			if errDDLPlayersInMatch != nil {
				panic(errDDLPlayersInMatch.Error())
			}
			dotDDLPlayer, errDDLPlayer := dotsql.LoadFromFile("./sql/ddl/player.sql")
			if errDDLPlayer != nil {
				panic(errDDLPlayer.Error())
			}
			dotDDLSoccerMatch, errDDLSoccerMatch := dotsql.LoadFromFile("./sql/ddl/soccer_match.sql")
			if errDDLSoccerMatch != nil {
				panic(errDDLSoccerMatch.Error())
			}
			dotDDLWorldCup, errDDLWorldCup := dotsql.LoadFromFile("./sql/ddl/world_cup.sql")
			if errDDLWorldCup != nil {
				panic(errDDLWorldCup.Error())
			}
			dotFeatureTopWinning, errFeatureTopWinning := dotsql.LoadFromFile("./sql/features/top_winning_countries_first_place.sql")
			if errFeatureTopWinning != nil {
				panic(errFeatureTopWinning.Error())
			}
			dotFeaturePodium, errFeaturePodium := dotsql.LoadFromFile("./sql/features/top_winning_countries_podium.sql")
			if errFeaturePodium != nil {
				panic(errFeaturePodium.Error())
			}
			dotMostPopularWCAttendance, errMostPopularWCAttendance := dotsql.LoadFromFile("./sql/features/most_popular_WC_attendance.sql")
			if errMostPopularWCAttendance != nil {
				panic(errMostPopularWCAttendance.Error())
			}
			dotTopCountriesGoalScorer, errTopCountriesGoalScorer := dotsql.LoadFromFile("./sql/features/countries_top_goal_scorer.sql")
			if errTopCountriesGoalScorer != nil {
				panic(errTopCountriesGoalScorer.Error())
			}
			dotAverageGoalScored, errAverageGoalScored := dotsql.LoadFromFile("./sql/features/countries_average_goals_scored_match.sql")
			if errAverageGoalScored != nil {
				panic(errAverageGoalScored.Error())
			}

			dotPopulateWorldCupMatches, errPopulateWorldCupMatches := dotsql.LoadFromFile("./sql/prod_data/csv/WorldCupMatchesCSV.sql")
			if errPopulateWorldCupMatches != nil {
				panic(errPopulateWorldCupMatches.Error())
			}

			dotPopulateWorldCupPlayers, errPopulateWorldCupPlayers := dotsql.LoadFromFile("./sql/prod_data/csv/WorldCupPlayersCSV.sql")
			if errPopulateWorldCupPlayers != nil {
				panic(errPopulateWorldCupMatches.Error())
			}

			dotPopulateWorldCup, errPopulateWorldCup := dotsql.LoadFromFile("./sql/prod_data/csv/WorldCupsCSV.sql")
			if errPopulateWorldCup != nil {
				panic(errPopulateWorldCup.Error())
			}

			dotPopulateProdCountry, errPopulateProdCountry := dotsql.LoadFromFile("./sql/prod_data/populate/country.sql")
			if errPopulateProdCountry != nil {
				panic(errPopulateProdCountry.Error())
			}
			dotPopulateProdPlayerInMatch, errPopulateProdPlayerInMatch := dotsql.LoadFromFile("./sql/prod_data/populate/player_plays_in_match.sql")
			if errPopulateProdPlayerInMatch != nil {
				panic(errPopulateProdPlayerInMatch.Error())
			}
			dotPopulateProdPlayer, errPopulateProdPlayer := dotsql.LoadFromFile("./sql/prod_data/populate/player.sql")
			if errPopulateProdPlayer != nil {
				panic(errPopulateProdPlayer.Error())
			}
			dotPopulateMatch, errPopulateMatch := dotsql.LoadFromFile("./sql/prod_data/populate/soccer_match.sql")
			if errPopulateMatch != nil {
				panic(errPopulateMatch.Error())
			}
			dotPopulateWC, errPopulateWC := dotsql.LoadFromFile("./sql/prod_data/populate/world_cup.sql")
			if errPopulateWC != nil {
				panic(errPopulateWC.Error())
			}

			dotWinRatio, errWinRatio := dotsql.LoadFromFile("./sql/features/countries_win_ratio.sql")
			if errWinRatio != nil {
				panic(errPopulateWC.Error())
			}

			dotPlayerMostStarted, errPlayerMostStarted := dotsql.LoadFromFile("./sql/features/players_most_started.sql")
			if errPlayerMostStarted != nil {
				panic(errPlayerMostStarted.Error())
			}

			dotDDLUser, errDDLUser := dotsql.LoadFromFile("./sql/ddl/user.sql")
			if errDDLUser != nil {
				panic(errDDLUser.Error())
			}

			dotUserQueries, errUserQueries := dotsql.LoadFromFile("./sql/queries/userQueries.sql")
			if errUserQueries != nil {
				panic(errUserQueries.Error())
			}

			println("Merging")
			dot := dotsql.Merge(
				dotPlayerMostStarted,
				dotPopulateProdPlayerInMatch,
				dotPopulateProdPlayer,
				dotPopulateWC,
				dotPopulateMatch,
				dotDDLCountries,
				dotDDLPlayersInMatch,
				dotDDLPlayer,
				dotDDLSoccerMatch,
				dotDDLWorldCup,
				dotFeatureTopWinning,
				dotPopulateProdCountry,
				dotFeaturePodium,
				dotMostPopularWCAttendance,
				dotTopCountriesGoalScorer,
				dotAverageGoalScored,
				dotQueryCountry,
				dotPopulateWorldCupMatches,
				dotPopulateWorldCupPlayers,
				dotPopulateWorldCup,
				dotWinRatio,
				dotDDLUser,
				dotUserQueries,
			)
			println("Done Loading")

			dot.Exec(db, "ddl-country-set-foreign-0")
			dot.Exec(db, "ddl-country-drop-if-exists")
			dot.Exec(db, "ddl-country-set-foreign-1")
			_, errCreateCountries := dot.Exec(db, "ddl-country-create-country")
			if errCreateCountries != nil {
				panic(errCreateCountries.Error())
			}

			dot.Exec(db, "ddl-player-set-foreign-0")
			dot.Exec(db, "ddl-player-drop-if-exists")
			dot.Exec(db, "ddl-player-set-foreign-1")
			_, errCreatePlayer := dot.Exec(db, "ddl-player-create-player")
			if errCreatePlayer != nil {
				panic(errCreatePlayer.Error())
			}

			dot.Exec(db, "ddl-world-cup-set-foreign-0")
			dot.Exec(db, "ddl-world-cup-drop-if-exists")
			dot.Exec(db, "ddl-world-cup-set-foreign-1")
			_, errCreateWorldCup := dot.Exec(db, "ddl-world-cup-create-table")
			if errCreateWorldCup != nil {
				panic(errCreateWorldCup.Error())
			}

			dot.Exec(db, "ddl-match-set-foreign-0")
			dot.Exec(db, "ddl-match-drop-if-exists")
			dot.Exec(db, "ddl-match-set-foreign-1")
			_, errCreateMatches := dot.Exec(db, "ddl-match-create-matches")
			if errCreateMatches != nil {
				panic(errCreateMatches.Error())
			}

			dot.Exec(db, "ddl-player-play-set-foreign-0")
			dot.Exec(db, "ddl-player-drop-table-if-exists")
			dot.Exec(db, "ddl-player-play-set-foreign-1")
			_, errCreatePlayerPlay := dot.Exec(db, "ddl-player-play-create-table")
			if errCreatePlayerPlay != nil {
				panic(errCreatePlayerPlay.Error())
			}

			dot.Exec(db, "ddl-user-set-foreign-1")
			_, errCreateUser := dot.Exec(db, "ddl-user-create-user")
			if errCreateUser != nil {
				panic(errCreateUser.Error())
			}

			var wg sync.WaitGroup
			wg.Add(1)
			csvMatches(&wg, db, dot)
			wg.Add(1)
			csvPlayers(&wg, db, dot)
			wg.Add(1)
			csvWorldCups(&wg, db, dot)
			wg.Wait()

			dot.Exec(db, "prod-data-country")
			dot.Exec(db, "prod-data-player")
			dot.Exec(db, "prod-data-world-cup")
			dot.Exec(db, "prod-data-match")
			dot.Exec(db, "prod-data-player-in-match")

			dot.Exec(db, "View-Country-Goal-Per-Match")
			dot.Exec(db, "View-Country-Winning-Matches")

			dot.Exec(db, "create-view-CountryGoalsScored")
			dot.Exec(db, "create-view-CountryTotalAppearances")
			dot.Exec(db, "create-view-CountryAvgData")

			dot.Exec(db, "create-view-CountryGoalScorers")
			dot.Exec(db, "create-view-CountryMostGoals")

			dot.Exec(db, "create-view-WorldCupAttendances")

			dot.Exec(db, "view-player-games-started")
			dot.Exec(db, "view-games-played")

			dot.Exec(db, "top-winning-countries-first-view")

			dot.Exec(db, "top-winning-countries-first-view-podium")
			dot.Exec(db, "top-winning-countries-second-view-podium")
			dot.Exec(db, "top-winning-countries-third-view-podium")

			initalizedDOT = dot
		} else {
			fmt.Println("Single instance already created.")
		}
	}
	return initalizedDOT
}
