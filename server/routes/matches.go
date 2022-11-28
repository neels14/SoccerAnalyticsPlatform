package routes

import (
	"backend/initalize"
	"backend/types"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllMatches(c *gin.Context) {
	dot := initalize.GetInstance()
	limit, errLimit := strconv.Atoi(c.Query("limit"))
	pageNumber, errPage := strconv.Atoi(c.Query("page"))
	if errLimit != nil {
		limit = 10
	}
	if errPage != nil {
		pageNumber = 1
	}
	rows, errAllMatches := dot.Query(initalize.GetDB(), "get-all-matches", limit, (pageNumber-1)*limit)
	if errAllMatches != nil {
		panic(errAllMatches.Error())
	}
	defer rows.Close()
	matches := []types.SoccerMatchResponse{}
	for rows.Next() {
		var match types.SoccerMatchResponse
		err := rows.Scan(&match.Year, &match.Stage, &match.HomeTeam, &match.AwayTeam, &match.Date, &match.City, &match.Stadium, &match.Attendance)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		matches = append(matches, match)

	}
	rows, errCountryCount := dot.Query(initalize.GetDB(), "count-matches")
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

	c.JSON(200, gin.H{"data": matches, "num_page": math.Ceil(float64(num_row) / float64(limit))})

}
