package routes

import (
	"backend/initalize"
	"backend/types"

	"github.com/gin-gonic/gin"
)

func GetAllPlayer(c *gin.Context) {
	dot := initalize.GetInstance()
	rows, errAllPlayer := dot.Query(initalize.GetDB(), "get-all-players")
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

	c.JSON(200, players)

}

func GetTopScorers(c *gin.Context) {
	dot := initalize.GetInstance()
	rows, errAllPlayer := dot.Query(initalize.GetDB(), "create-table-TopGoalScorerByCountry-without-country")
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

	c.JSON(200, players)

}

func GetTopScorersWithCountry(c *gin.Context) {
	country := c.Param("country")
	dot := initalize.GetInstance()
	rows, errAllPlayer := dot.Query(initalize.GetDB(), "create-table-TopGoalScorerByCountry", country)
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

	c.JSON(200, players)

}
