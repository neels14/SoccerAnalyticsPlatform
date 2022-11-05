package routes

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllMatches(c *gin.Context) {
	c.JSON(200, gin.H{
		"matches": "all",
	})
}

func GetMatchByYear(c *gin.Context) {
	year, err := strconv.Atoi(c.Param("year"))
	if err != nil {
		panic(err)
	}
	c.JSON(200, year)
}
