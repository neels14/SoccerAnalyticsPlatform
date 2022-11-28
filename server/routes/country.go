package routes

import (
	"backend/initalize"
	"backend/types"
	"fmt"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllCountries(c *gin.Context) {
	dot := initalize.GetInstance()
	limit, errLimit := strconv.Atoi(c.Query("limit"))
	pageNumber, errPage := strconv.Atoi(c.Query("page"))
	if errLimit != nil {
		fmt.Println("Limit")
		limit = 10
	}
	if errPage != nil {
		fmt.Println("Page")
		pageNumber = 1
	}
	rows, errAllCountries := dot.Query(initalize.GetDB(), "get-all-countries", limit, (pageNumber-1)*limit)
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
	rows, errCountryCount := dot.Query(initalize.GetDB(), "count-countries")
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

	c.JSON(200, gin.H{"data": countries, "num_page": math.Ceil(float64(num_row) / float64(limit))})
}

func AverageGoalByCountry(c *gin.Context) {
	dot := initalize.GetInstance()
	limit, errLimit := strconv.Atoi(c.Query("limit"))
	pageNumber, errPage := strconv.Atoi(c.Query("page"))
	if errLimit != nil {
		limit = 10
	}
	if errPage != nil {
		pageNumber = 1
	}
	rows, errAllCountries := dot.Query(initalize.GetDB(), "create-table-AvgGoalsInMatchByCountry-without-country", limit, (pageNumber-1)*limit)
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
	rows, errCountryCount := dot.Query(initalize.GetDB(), "count-create-table-AvgGoalsInMatchByCountry-without-country")
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

	c.JSON(200, gin.H{"data": countries, "num_page": math.Ceil(float64(num_row) / float64(limit))})
}

func AverageGoalForCountry(c *gin.Context) {
	dot := initalize.GetInstance()
	name := c.Param("name")
	limit, errLimit := strconv.Atoi(c.Query("limit"))
	pageNumber, errPage := strconv.Atoi(c.Query("page"))
	if errLimit != nil {
		limit = 10
	}
	if errPage != nil {
		pageNumber = 1
	}
	rows, errAllCountries := dot.Query(initalize.GetDB(), "create-table-AvgGoalsInMatchByCountry", name, limit, (pageNumber-1)*limit)
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
	rows, errCountryCount := dot.Query(initalize.GetDB(), "count-create-table-AvgGoalsInMatchByCountry", name)
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

	c.JSON(200, gin.H{"data": countries, "num_page": math.Ceil(float64(num_row) / float64(limit))})

}

func WinRatio(c *gin.Context) {
	dot := initalize.GetInstance()
	limit, errLimit := strconv.Atoi(c.Query("limit"))
	pageNumber, errPage := strconv.Atoi(c.Query("page"))
	if errLimit != nil {
		limit = 10
	}
	if errPage != nil {
		pageNumber = 1
	}
	rows, errAllCountries := dot.Query(initalize.GetDB(), "Country-Win-Ratio", limit, (pageNumber-1)*limit)
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
	rows, errCountryCount := dot.Query(initalize.GetDB(), "Count-Country-Win-Ratio")
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

	c.JSON(200, gin.H{"data": countries, "num_page": math.Ceil(float64(num_row) / float64(limit))})

}

func WinRatioWithCountry(c *gin.Context) {
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
	rows, errAllCountries := dot.Query(initalize.GetDB(), "Country-Win-Ratio-specific-country", country, limit, (pageNumber-1)*limit)
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
	rows, errCountryCount := dot.Query(initalize.GetDB(), "Count-Country-Win-Ratio-specific-country", country)
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

	c.JSON(200, gin.H{"data": countries, "num_page": math.Ceil(float64(num_row) / float64(limit))})

}
