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

func GetCountry(c *gin.Context) {
	name := c.Param("name")
	c.JSON(200, name)
}
