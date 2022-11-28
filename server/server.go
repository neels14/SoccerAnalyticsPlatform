package main

import (
	"backend/initalize"
	"backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	initalize.GetInstance()

	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		country := v1.Group("/country")
		{
			country.GET("/", routes.GetAllCountries)
			country.GET("/averageGoal", routes.AverageGoalByCountry)
			country.GET("/averageGoal/:name", routes.AverageGoalForCountry)
			country.GET("/winRatio", routes.WinRatio)
			country.GET("/winRatio/:country", routes.WinRatioWithCountry)
		}

		matches := v1.Group("/matches")
		{
			matches.GET("/", routes.GetAllMatches)
		}
		players := v1.Group("/players")
		{
			players.GET("/", routes.GetAllPlayer)
			players.GET("/topScorer", routes.GetTopScorers)
			players.GET("/topScorer/:country", routes.GetTopScorersWithCountry)
			players.GET("/mostStarted", routes.GetMostStarted)
			players.GET("/mostStarted/:country", routes.GetMostStartedWithCountry)
		}
		world_cup := v1.Group("/worldCup")
		{
			world_cup.GET("/", routes.GetAllWorldCups)
			world_cup.GET("/popular", routes.MostPopular)
			world_cup.GET("/first", routes.First)
			world_cup.GET("/first/:country", routes.FirstWithName)
			world_cup.GET("/podium", routes.Podium)
			world_cup.GET("/podium/:country", routes.PodiumWithCountry)
		}
		user := v1.Group("/user")
		{
			user.GET("/:username", routes.GetUserByUsername)
			user.GET("/:username/:password", routes.VerifyUser)
			user.POST("/", routes.AddUser)
		}
	}
	r.Run()
	initalize.GetDB().Close()
}
