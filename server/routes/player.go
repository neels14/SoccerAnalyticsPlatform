package routes

import (
	"backend/initalize"
	"backend/types"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllPlayer(c *gin.Context) {
	dot := initalize.GetInstance()
	limit, errLimit := strconv.Atoi(c.Query("limit"))
	pageNumber, errPage := strconv.Atoi(c.Query("page"))
	if errLimit != nil {
		limit = 10
	}
	if errPage != nil {
		pageNumber = 1
	}
	rows, errAllPlayer := dot.Query(initalize.GetDB(), "get-all-players", limit, (pageNumber-1)*limit)
	if errAllPlayer != nil {
		panic(errAllPlayer.Error())
	}
	defer rows.Close()
	players := []types.PlayerResponse{}
	for rows.Next() {
		var player types.PlayerResponse
		err := rows.Scan(&player.Name, &player.ShirtNumber, &player.Country)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		players = append(players, player)

	}
	rows, errCountryCount := dot.Query(initalize.GetDB(), "count-players")
	if errCountryCount != nil {
		panic(errCountryCount.Error())
	}
	defer rows.Close()
	var num_row int
	for rows.Next() {
		err := rows.Scan(&num_row)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

	}

	c.JSON(200, gin.H{"data": players, "num_page": math.Ceil(float64(num_row) / float64(limit))})

}

func GetTopScorers(c *gin.Context) {
	limit, errLimit := strconv.Atoi(c.Query("limit"))
	pageNumber, errPage := strconv.Atoi(c.Query("page"))
	if errLimit != nil {
		limit = 10
	}
	if errPage != nil {
		pageNumber = 1
	}
	dot := initalize.GetInstance()
	rows, errAllPlayer := dot.Query(initalize.GetDB(), "create-table-TopGoalScorerByCountry-without-country", limit, (pageNumber-1)*limit)
	if errAllPlayer != nil {
		panic(errAllPlayer.Error())
	}
	defer rows.Close()
	players := []types.TopScorer{}
	for rows.Next() {
		var player types.TopScorer
		err := rows.Scan(&player.Name, &player.TopScorer, &player.Goals)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		players = append(players, player)

	}
	rows, errCountryCount := dot.Query(initalize.GetDB(), "count-create-table-TopGoalScorerByCountry-without-country")
	if errCountryCount != nil {
		panic(errCountryCount.Error())
	}
	defer rows.Close()
	var num_row int
	for rows.Next() {
		err := rows.Scan(&num_row)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

	}

	c.JSON(200, gin.H{"data": players, "num_page": math.Ceil(float64(num_row) / float64(limit))})

}

func GetTopScorersWithCountry(c *gin.Context) {
	country := c.Param("country")
	limit, errLimit := strconv.Atoi(c.Query("limit"))
	pageNumber, errPage := strconv.Atoi(c.Query("page"))
	if errLimit != nil {
		limit = 10
	}
	if errPage != nil {
		pageNumber = 1
	}
	dot := initalize.GetInstance()
	rows, errAllPlayer := dot.Query(initalize.GetDB(), "create-table-TopGoalScorerByCountry", country, limit, (pageNumber-1)*limit)
	if errAllPlayer != nil {
		panic(errAllPlayer.Error())
	}
	defer rows.Close()
	players := []types.TopScorer{}
	for rows.Next() {
		var player types.TopScorer
		err := rows.Scan(&player.Name, &player.TopScorer, &player.Goals)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		players = append(players, player)

	}
	rows, errCountryCount := dot.Query(initalize.GetDB(), "count-create-table-TopGoalScorerByCountry", country)
	if errCountryCount != nil {
		panic(errCountryCount.Error())
	}
	defer rows.Close()
	var num_row int
	for rows.Next() {
		err := rows.Scan(&num_row)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

	}

	c.JSON(200, gin.H{"data": players, "num_page": math.Ceil(float64(num_row) / float64(limit))})

}

func GetMostStarted(c *gin.Context) {
	limit, errLimit := strconv.Atoi(c.Query("limit"))
	pageNumber, errPage := strconv.Atoi(c.Query("page"))
	if errLimit != nil {
		limit = 10
	}
	if errPage != nil {
		pageNumber = 1
	}
	dot := initalize.GetInstance()
	rows, errAllPlayer := dot.Query(initalize.GetDB(), "player-most-started", limit, (pageNumber-1)*limit)
	if errAllPlayer != nil {
		panic(errAllPlayer.Error())
	}
	defer rows.Close()
	players := []types.PlayerMostStarted{}
	for rows.Next() {
		var player types.PlayerMostStarted
		err := rows.Scan(&player.Name, &player.ShirtNumber, &player.Country, &player.StartRatio)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		players = append(players, player)

	}
	rows, errCountryCount := dot.Query(initalize.GetDB(), "count-player-most-started")
	if errCountryCount != nil {
		panic(errCountryCount.Error())
	}
	defer rows.Close()
	var num_row int
	for rows.Next() {
		err := rows.Scan(&num_row)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

	}

	c.JSON(200, gin.H{"data": players, "num_page": math.Ceil(float64(num_row) / float64(limit))})
}

func GetMostStartedWithCountry(c *gin.Context) {
	limit, errLimit := strconv.Atoi(c.Query("limit"))
	pageNumber, errPage := strconv.Atoi(c.Query("page"))
	if errLimit != nil {
		limit = 10
	}
	if errPage != nil {
		pageNumber = 1
	}
	country := c.Param("country")
	dot := initalize.GetInstance()
	rows, errAllPlayer := dot.Query(initalize.GetDB(), "player-most-started-specific-country", country, limit, (pageNumber-1)*limit)
	if errAllPlayer != nil {
		panic(errAllPlayer.Error())
	}
	defer rows.Close()
	players := []types.PlayerMostStarted{}
	for rows.Next() {
		var player types.PlayerMostStarted
		err := rows.Scan(&player.Name, &player.ShirtNumber, &player.Country, &player.StartRatio)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		players = append(players, player)

	}
	rows, errCountryCount := dot.Query(initalize.GetDB(), "count-player-most-started-specific-country")
	if errCountryCount != nil {
		panic(errCountryCount.Error())
	}
	defer rows.Close()
	var num_row int
	for rows.Next() {
		err := rows.Scan(&num_row)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

	}

	c.JSON(200, gin.H{"data": players, "num_page": math.Ceil(float64(num_row) / float64(limit))})
}
