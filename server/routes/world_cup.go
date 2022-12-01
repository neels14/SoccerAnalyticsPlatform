package routes

import (
	"backend/initalize"
	"backend/types"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllWorldCups(c *gin.Context) {
	limit, errLimit := strconv.Atoi(c.Query("limit"))
	pageNumber, errPage := strconv.Atoi(c.Query("page"))
	if errLimit != nil {
		limit = 10
	}
	if errPage != nil {
		pageNumber = 1
	}
	dot := initalize.GetInstance()
	rows, errAllWC := dot.Query(initalize.GetDB(), "get-all-wc", limit, (pageNumber-1)*limit)
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
	rows, errCountryCount := dot.Query(initalize.GetDB(), "count-wc")
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

	c.JSON(200, gin.H{"data": wcs, "num_page": math.Ceil(float64(num_row) / float64(limit))})

}

func MostPopular(c *gin.Context) {
	limit, errLimit := strconv.Atoi(c.Query("limit"))
	pageNumber, errPage := strconv.Atoi(c.Query("page"))
	if errLimit != nil {
		limit = 10
	}
	if errPage != nil {
		pageNumber = 1
	}
	dot := initalize.GetInstance()
	rows, errAllWC := dot.Query(initalize.GetDB(), "create-table-MostPopularWCByAttendance", limit, (pageNumber-1)*limit)
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
	rows, errCountryCount := dot.Query(initalize.GetDB(), "count-create-table-MostPopularWCByAttendance")
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

	c.JSON(200, gin.H{"data": wcs, "num_page": math.Ceil(float64(num_row) / float64(limit))})

}

func First(c *gin.Context) {
	limit, errLimit := strconv.Atoi(c.Query("limit"))
	pageNumber, errPage := strconv.Atoi(c.Query("page"))
	if errLimit != nil {
		limit = 10
	}
	if errPage != nil {
		pageNumber = 1
	}
	dot := initalize.GetInstance()
	rows, errAllWC := dot.Query(initalize.GetDB(), "top-winning-countries-first-without-name", limit, (pageNumber-1)*limit)
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
	rows, errCountryCount := dot.Query(initalize.GetDB(), "count-top-winning-countries-first-without-name")
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

	c.JSON(200, gin.H{"data": wcs, "num_page": math.Ceil(float64(num_row) / float64(limit))})

}

func FirstWithName(c *gin.Context) {
	limit, errLimit := strconv.Atoi(c.Query("limit"))
	pageNumber, errPage := strconv.Atoi(c.Query("page"))
	if errLimit != nil {
		limit = 10
	}
	if errPage != nil {
		pageNumber = 1
	}
	dot := initalize.GetInstance()
	country := c.Param("country")
	rows, errAllWC := dot.Query(initalize.GetDB(), "top-winning-countries-first", country, limit, (pageNumber-1)*limit)
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
	rows, errCountryCount := dot.Query(initalize.GetDB(), "count-top-winning-countries-first", country)
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

	c.JSON(200, gin.H{"data": wcs, "num_page": math.Ceil(float64(num_row) / float64(limit))})

}

func Podium(c *gin.Context) {
	limit, errLimit := strconv.Atoi(c.Query("limit"))
	pageNumber, errPage := strconv.Atoi(c.Query("page"))
	if errLimit != nil {
		limit = 10
	}
	if errPage != nil {
		pageNumber = 1
	}
	dot := initalize.GetInstance()
	rows, errAllWC := dot.Query(initalize.GetDB(), "top-winning-countries-podium", limit, (pageNumber-1)*limit)
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
	rows, errCountryCount := dot.Query(initalize.GetDB(), "count-top-winning-countries-podium")
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

	c.JSON(200, gin.H{"data": wcs, "num_page": math.Ceil(float64(num_row) / float64(limit))})

}

func PodiumWithCountry(c *gin.Context) {
	limit, errLimit := strconv.Atoi(c.Query("limit"))
	pageNumber, errPage := strconv.Atoi(c.Query("page"))
	if errLimit != nil {
		limit = 10
	}
	if errPage != nil {
		pageNumber = 1
	}
	dot := initalize.GetInstance()
	country := c.Param("country")
	rows, errAllWC := dot.Query(initalize.GetDB(), "top-winning-countries-podium-with-country-name", country, limit, (pageNumber-1)*limit)
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
	rows, errCountryCount := dot.Query(initalize.GetDB(), "count-top-winning-countries-podium-with-country-name", country)
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

	c.JSON(200, gin.H{"data": wcs, "num_page": math.Ceil(float64(num_row) / float64(limit))})

}
