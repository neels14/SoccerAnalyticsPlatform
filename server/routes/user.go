package routes

import (
	"backend/initalize"
	"backend/types"
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserByUsername(c *gin.Context) {
	dot := initalize.GetInstance()
	username := c.Param("username")
	rows, errUser := dot.Query(initalize.GetDB(), "get-user", username)
	if errUser != nil {
		panic(errUser.Error())
	}
	defer rows.Close()
	var users *types.User = nil
	for rows.Next() {
		var user types.User
		var password string
		err := rows.Scan(&user.Username, &user.FirstName, &user.LastName, &user.Email, &password)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		users = &user
	}

	c.JSON(200, users)
}

func VerifyUser(c *gin.Context) {
	dot := initalize.GetInstance()
	username := c.Param("username")
	inputPassword := c.Param("password")
	rows, errUser := dot.Query(initalize.GetDB(), "get-user", username)
	if errUser != nil {
		panic(errUser.Error())
	}
	defer rows.Close()
	foundUser := false
	correctPassword := false
	for rows.Next() {
		foundUser = true
		var user types.User
		var dbPassword string
		err := rows.Scan(&user.Username, &user.FirstName, &user.LastName, &user.Email, &dbPassword)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// Compare the inputted password with user db password
		hashedInput := md5.Sum([]byte(inputPassword))
		hashedInputStr := hex.EncodeToString(hashedInput[:])
		if hashedInputStr == dbPassword {
			correctPassword = true
		}
	}

	var returnUserValid *types.UserValid = nil
	if foundUser && correctPassword {
		returnUserValid = &types.UserValid{Valid: true}
	} else {
		returnUserValid = &types.UserValid{Valid: false}
	}

	c.JSON(200, returnUserValid)
}

func AddUser(c *gin.Context) {
	dot := initalize.GetInstance()
	var newUser types.NewUser
	err := c.BindJSON(&newUser)
	if err != nil {
		panic(err.Error())
	}
	_, errUser := dot.Query(initalize.GetDB(), "insert-user", newUser.Username, newUser.FirstName, newUser.LastName, newUser.Email, newUser.Password)
	if errUser != nil {
		// Username already exists or email is already taken.
		c.JSON(400, nil)
		return
	}
	c.JSON(http.StatusCreated, types.User{
		Username:  newUser.Username,
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Email:     newUser.Email,
	})
}
