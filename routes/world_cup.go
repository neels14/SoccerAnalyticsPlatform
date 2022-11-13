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
