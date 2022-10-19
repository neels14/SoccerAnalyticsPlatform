package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/qustavo/dotsql"
)

const testDBName string = "testdb"
const prodDBName string = "proddb"

func main() {
	// Open up our database connection.
	// The database is called testdb
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/"+testDBName)

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	//Loads queries from file
	dot := SetupDot(db)

	fmt.Println("---Countries that has finished at the top the most---")
	featureTop(dot, db)
	fmt.Println("---Countries that has placed the most---")
	featurePodium(dot, db)
	fmt.Println("---WC that have the most popular attendance---")
	featureMostPopularWCAttendance(dot, db)
	fmt.Println("---Highest Goal Scorer per country---")
	featureCountriesTopGoalScorer(dot, db)
	fmt.Println("---Average goals scored per countries---")
	featureCountriesAverageGoalScoredMatch(dot, db)

}

func featureMostPopularWCAttendance(dot *dotsql.DotSql, db *sql.DB) {
	dot.Exec(db, "create-view-WorldCupAttendances")
	rows, errRunFeaturePodium := dot.Query(db, "create-table-MostPopularWCByAttendance")
	if errRunFeaturePodium != nil {
		panic(errRunFeaturePodium.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var year int
		var attendance int
		if err := rows.Scan(&year, &attendance); err != nil {
			panic(err)
		}
		fmt.Println("Year:", year, "has attendence", attendance)
	}
}

func featureCountriesTopGoalScorer(dot *dotsql.DotSql, db *sql.DB) {
	dot.Exec(db, "create-view-CountryGoalScorers")
	dot.Exec(db, "create-view-CountryMostGoals")
	rows, errRunFeaturePodium := dot.Query(db, "create-table-TopGoalScorerByCountry")
	if errRunFeaturePodium != nil {
		panic(errRunFeaturePodium.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var country_name string
		var goal_name string
		var goals int
		if err := rows.Scan(&country_name, &goal_name, &goals); err != nil {
			panic(err)
		}
		fmt.Println(country_name, goal_name, goals)
	}
}

func featureCountriesAverageGoalScoredMatch(dot *dotsql.DotSql, db *sql.DB) {
	dot.Exec(db, "create-view-CountryGoalsScored")
	dot.Exec(db, "create-view-CountryTotalAppearances")
	dot.Exec(db, "create-view-CountryAvgData")
	rows, errRunFeaturePodium := dot.Query(db, "create-table-AvgGoalsInMatchByCountry")
	if errRunFeaturePodium != nil {
		panic(errRunFeaturePodium.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var country_name string
		var avg_goals_match float32
		if err := rows.Scan(&country_name, &avg_goals_match); err != nil {
			panic(err)
		}
		fmt.Println(country_name, ":", avg_goals_match)
	}
}

func featurePodium(dot *dotsql.DotSql, db *sql.DB) {
	dot.Exec(db, "top-winning-countries-first-view-podium")
	dot.Exec(db, "top-winning-countries-second-view-podium")
	dot.Exec(db, "top-winning-countries-third-view-podium")
	rows, errRunFeaturePodium := dot.Query(db, "top-winning-countries-podium")
	if errRunFeaturePodium != nil {
		panic(errRunFeaturePodium.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var country_name string
		var timesPlaced int
		if err := rows.Scan(&country_name, &timesPlaced); err != nil {
			panic(err)
		}
		fmt.Println(country_name, "placed", timesPlaced, "times")
	}
}
func featureTop(dot *dotsql.DotSql, db *sql.DB) {
	dot.Exec(db, "top-winning-countries-first-view")
	rows, errRunFeatureTopWinning := dot.Query(db, "top-winning-countries-first")
	if errRunFeatureTopWinning != nil {
		panic(errRunFeatureTopWinning.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var country_name string
		var wins int
		if err := rows.Scan(&country_name, &wins); err != nil {
			panic(err)
		}
		fmt.Println(country_name, "won", wins, "times")
	}
}

func SetupDot(db *sql.DB) *dotsql.DotSql {
	//Loads queries from file
	dotDDLCountries, errDDLCountries := dotsql.LoadFromFile("../sql/ddl/country.sql")
	if errDDLCountries != nil {
		panic(errDDLCountries.Error())
	}
	dotDDLPlayersInMatch, errDDLPlayersInMatch := dotsql.LoadFromFile("../sql/ddl/player_plays_in_match.sql")
	if errDDLPlayersInMatch != nil {
		panic(errDDLCountries.Error())
	}
	dotDDLPlayer, errDDLPlayer := dotsql.LoadFromFile("../sql/ddl/player.sql")
	if errDDLPlayer != nil {
		panic(errDDLCountries.Error())
	}
	dotDDLSoccerMatch, errDDLSoccerMatch := dotsql.LoadFromFile("../sql/ddl/soccer_match.sql")
	if errDDLSoccerMatch != nil {
		panic(errDDLSoccerMatch.Error())
	}
	dotDDLWorldCup, errDDLWorldCup := dotsql.LoadFromFile("../sql/ddl/world_cup.sql")
	if errDDLWorldCup != nil {
		panic(errDDLWorldCup.Error())
	}
	dotFeatureTopWinning, errFeatureTopWinning := dotsql.LoadFromFile("../sql/features/top_winning_countries_first_place.sql")
	if errFeatureTopWinning != nil {
		panic(errFeatureTopWinning.Error())
	}
	dotFeaturePodium, errFeaturePodium := dotsql.LoadFromFile("../sql/features/top_winning_countries_podium.sql")
	if errFeaturePodium != nil {
		panic(errFeaturePodium.Error())
	}
	dotMostPopularWCAttendance, errMostPopularWCAttendance := dotsql.LoadFromFile("../sql/features/most_popular_WC_attendance.sql")
	if errMostPopularWCAttendance != nil {
		panic(errMostPopularWCAttendance.Error())
	}
	dotTopCountriesGoalScorer, errTopCountriesGoalScorer := dotsql.LoadFromFile("../sql/features/countries_top_goal_scorer.sql")
	if errTopCountriesGoalScorer != nil {
		panic(errTopCountriesGoalScorer.Error())
	}
	dotAverageGoalScored, errAverageGoalScored := dotsql.LoadFromFile("../sql/features/countries_average_goals_scored_match.sql")
	if errAverageGoalScored != nil {
		panic(errAverageGoalScored.Error())
	}
	dotPopulateTestData, errPopulateTestData := dotsql.LoadFromFile("../sql/test_data/populate_test_db.sql")
	if errPopulateTestData != nil {
		panic(errPopulateTestData.Error())
	}
	dot := dotsql.Merge(dotDDLCountries, dotDDLPlayersInMatch, dotDDLPlayer, dotDDLSoccerMatch, dotDDLWorldCup, dotFeatureTopWinning, dotPopulateTestData, dotFeaturePodium, dotMostPopularWCAttendance, dotTopCountriesGoalScorer, dotAverageGoalScored)
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
	dot.Exec(db, "test-data-into-countries")
	dot.Exec(db, "test-data-into-players")
	dot.Exec(db, "test-data-into-world-cup")
	dot.Exec(db, "test-data-into-match")
	dot.Exec(db, "test-data-into-player-in-match")
	return dot
}
