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
			country.GET("/:name", routes.GetCountry)
		}

		matches := v1.Group("/matches")
		{
			matches.GET("/", routes.GetAllMatches)
			matches.GET("/:year", routes.GetMatchByYear)
		}

	}
	r.Run()
	initalize.GetDB().Close()
}
