package routes

import (
	"backend/initalize"
	"backend/types"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllMatches(c *gin.Context) {
	dot := initalize.GetInstance()
	rows, errAllMatches := dot.Query(initalize.GetDB(), "get-all-matches")
	if errAllMatches != nil {
		panic(errAllMatches.Error())
	}
	defer rows.Close()
	matches := []types.SoccerMatchResponse{}
	for rows.Next() {
		var match types.SoccerMatchResponse
		err := rows.Scan(&match.Year, &match.Stage, &match.HomeTeam, &match.AwayTeam, &match.Date, &match.City, &match.Stadium, &match.Stadium, &match.HomeCoach, &match.AwayCoach)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		matches = append(matches, match)

	}

	c.JSON(200, matches)

}

func GetMatchByYear(c *gin.Context) {
	year, err := strconv.Atoi(c.Param("year"))
	if err != nil {
		panic(err)
	}
	c.JSON(200, year)
}
