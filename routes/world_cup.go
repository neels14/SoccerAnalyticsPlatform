package routes

import (
	"backend/initalize"
	"backend/types"

	"github.com/gin-gonic/gin"
)

func GetAllWorldCups(c *gin.Context) {
	dot := initalize.GetInstance()
	rows, errAllWC := dot.Query(initalize.GetDB(), "get-all-wc")
	if errAllWC != nil {
		panic(errAllWC.Error())
	}
	defer rows.Close()
	wcs := []types.WorldCupResponse{}
	for rows.Next() {
		var wc types.WorldCupResponse
		err := rows.Scan(&wc.Year, &wc.HostCountry, &wc.FirstPlace, &wc.SecondPlace, &wc.ThirdPlace)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		wcs = append(wcs, wc)

	}

	c.JSON(200, wcs)

}

func MostPopular(c *gin.Context) {
	dot := initalize.GetInstance()
	rows, errAllWC := dot.Query(initalize.GetDB(), "create-table-MostPopularWCByAttendance")
	if errAllWC != nil {
		panic(errAllWC.Error())
	}
	defer rows.Close()
	wcs := []types.WorldCupAttendence{}
	for rows.Next() {
		var wc types.WorldCupAttendence
		err := rows.Scan(&wc.Year, &wc.Attendance)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		wcs = append(wcs, wc)

	}

	c.JSON(200, wcs)

}

func First(c *gin.Context) {
	dot := initalize.GetInstance()
	rows, errAllWC := dot.Query(initalize.GetDB(), "top-winning-countries-first-without-name")
	if errAllWC != nil {
		panic(errAllWC.Error())
	}
	defer rows.Close()
	wcs := []types.Winner{}
	for rows.Next() {
		var wc types.Winner
		err := rows.Scan(&wc.Country, &wc.Wins)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		wcs = append(wcs, wc)

	}

	c.JSON(200, wcs)

}

func FirstWithName(c *gin.Context) {
	dot := initalize.GetInstance()
	country := c.Param("country")
	rows, errAllWC := dot.Query(initalize.GetDB(), "top-winning-countries-first", country)
	if errAllWC != nil {
		panic(errAllWC.Error())
	}
	defer rows.Close()
	wcs := []types.Winner{}
	for rows.Next() {
		var wc types.Winner
		err := rows.Scan(&wc.Country, &wc.Wins)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		wcs = append(wcs, wc)

	}

	c.JSON(200, wcs)

}

func Podium(c *gin.Context) {
	dot := initalize.GetInstance()
	rows, errAllWC := dot.Query(initalize.GetDB(), "top-winning-countries-podium")
	if errAllWC != nil {
		panic(errAllWC.Error())
	}
	defer rows.Close()
	wcs := []types.Winner{}
	for rows.Next() {
		var wc types.Winner
		err := rows.Scan(&wc.Country, &wc.Wins)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		wcs = append(wcs, wc)

	}

	c.JSON(200, wcs)

}

func PodiumWithCountry(c *gin.Context) {
	dot := initalize.GetInstance()
	country := c.Param("country")
	rows, errAllWC := dot.Query(initalize.GetDB(), "top-winning-countries-podium-with-country-name", country)
	if errAllWC != nil {
		panic(errAllWC.Error())
	}
	defer rows.Close()
	wcs := []types.Winner{}
	for rows.Next() {
		var wc types.Winner
		err := rows.Scan(&wc.Country, &wc.Wins)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		wcs = append(wcs, wc)

	}

	c.JSON(200, wcs)

}
