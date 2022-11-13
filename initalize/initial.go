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
				panic(errDDLCountries.Error())
			}
			dotDDLPlayer, errDDLPlayer := dotsql.LoadFromFile("./sql/ddl/player.sql")
			if errDDLPlayer != nil {
				panic(errDDLCountries.Error())
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
			dot := dotsql.Merge(
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
			)
			dot.Exec(db, "ddl-country-set-foreign-0")
			dot.Exec(db, "ddl-country-drop-if-exists")
			dot.Exec(db, "ddl-country-set-foreign-1")
			_, errCreateCountries := dot.Exec(db, "ddl-country-create-country")
			if errCreateCountries != nil {
				panic(errCreateCountries.Error())
			}
			dot.Exec(db, "ddl-player-set-foreign-0")
			dot.Exec(db, "ddl-player-drop-if-exists")
			dot.Exec(db, "ddl-player-play-create-table")
			_, errCreatePlayer := dot.Exec(db, "ddl-player-create-player")
			if errCreatePlayer != nil {
				panic(errCreatePlayer.Error())
			}
			dot.Exec(db, "ddl-match-set-foreign-0")
			dot.Exec(db, "ddl-match-drop-if-exists")
			dot.Exec(db, "ddl-match-set-foreign-1")
			_, errCreateMatches := dot.Exec(db, "ddl-match-create-matches")
			if errCreateMatches != nil {
				panic(errCreateMatches.Error())
			}
			dot.Exec(db, "ddl-player-play-set-foreign-0")
			dot.Exec(db, "ddl-player-play-drop-table-if-exists")
			dot.Exec(db, "ddl-player-play-create-table")
			_, errCreatePlayerPlay := dot.Exec(db, "ddl-player-play-create-table")
			if errCreatePlayerPlay != nil {
				panic(errCreatePlayerPlay.Error())
			}

			dot.Exec(db, "ddl-world-cup-set-foreign-0")
			dot.Exec(db, "ddl-world-cup-drop-if-exists")
			dot.Exec(db, "ddl-world-cup-create-table")
			_, errCreateWorldCup := dot.Exec(db, "ddl-world-cup-create-table")
			if errCreateWorldCup != nil {
				panic(errCreateWorldCup.Error())
			}
			dot.Exec(db, "CSV-DROP-WorldCupMatchesCSV")
			dot.Exec(db, "CSV-Table-WorldCupMatches")
			dot.Exec(db, "CSV-Table-WorldCupMatches-Insert")

			dot.Exec(db, "CSV-DROP-WorldCupPlayersCSV")
			dot.Exec(db, "CSV-Table-WorldCupPlayersCSV")
			dot.Exec(db, "CSV-Table-WorldCupPlayers-Insert")

			dot.Exec(db, "CSV-DROP-WorldCupsCSV")
			dot.Exec(db, "CSV-Table-WorldCupsCSV")
			dot.Exec(db, "CSV-Table-WorldCupsCSV-Insert")

			dot.Exec(db, "prod-data-country")
			dot.Exec(db, "prod-data-player")
			dot.Exec(db, "prod-data-world-cup")
			dot.Exec(db, "prod-data-match")
			dot.Exec(db, "prod-data-player-in-match")

			initalizedDOT = dot
		} else {
			fmt.Println("Single instance already created.")
		}
	}
	return initalizedDOT
}
