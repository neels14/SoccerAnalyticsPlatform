package routes

import (
	"backend/initalize"
	"backend/types"

	"github.com/gin-gonic/gin"
)

func GetAllCountries(c *gin.Context) {
	dot := initalize.GetInstance()
	rows, errAllCountries := dot.Query(initalize.GetDB(), "get-all-countries")
	if errAllCountries != nil {
		panic(errAllCountries.Error())
	}
	defer rows.Close()
	countries := []types.CountryResponse{}
	for rows.Next() {
		var country types.CountryResponse
		err := rows.Scan(&country.Name, &country.TeamInitial)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		countries = append(countries, country)

	}

	c.JSON(200, countries)
}

func AverageGoalByCountry(c *gin.Context) {
	dot := initalize.GetInstance()
	rows, errAllCountries := dot.Query(initalize.GetDB(), "create-table-AvgGoalsInMatchByCountry-without-country")
	if errAllCountries != nil {
		panic(errAllCountries.Error())
	}
	defer rows.Close()
	countries := []types.AverageGoalByCountry{}
	for rows.Next() {
		var country types.AverageGoalByCountry
		err := rows.Scan(&country.Name, &country.AverageGoal)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		countries = append(countries, country)

	}

	c.JSON(200, countries)
}

func AverageGoalForCountry(c *gin.Context) {
	dot := initalize.GetInstance()
	name := c.Param("name")
	rows, errAllCountries := dot.Query(initalize.GetDB(), "create-table-AvgGoalsInMatchByCountry", name)
	if errAllCountries != nil {
		panic(errAllCountries.Error())
	}
	defer rows.Close()
	var countries *types.AverageGoalByCountry = nil
	for rows.Next() {
		var country types.AverageGoalByCountry
		err := rows.Scan(&country.Name, &country.AverageGoal)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		countries = &country

	}
	c.JSON(200, countries)

}

func WinRatio(c *gin.Context) {
	dot := initalize.GetInstance()
	rows, errAllCountries := dot.Query(initalize.GetDB(), "Country-Win-Ratio")
	if errAllCountries != nil {
		panic(errAllCountries.Error())
	}
	defer rows.Close()
	countries := []types.WinRatio{}
	for rows.Next() {
		var country types.WinRatio
		err := rows.Scan(&country.Country, &country.WinRatio)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		countries = append(countries, country)

	}
	c.JSON(200, countries)

}

func WinRatioWithCountry(c *gin.Context) {
    country := c.Param("country")
	dot := initalize.GetInstance()
	rows, errAllCountries := dot.Query(initalize.GetDB(), "Country-Win-Ratio-specific-country", country)
	if errAllCountries != nil {
		panic(errAllCountries.Error())
	}
	defer rows.Close()
	var countries *types.WinRatio = nil
	for rows.Next() {
		var country types.WinRatio
		err := rows.Scan(&country.Country, &country.WinRatio)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		countries = &country

	}
	c.JSON(200, countries)

}
