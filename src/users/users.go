package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Users(c *gin.Context) {
	var id string = c.Query("id")
	if id != "" {
		user, err := RetrieveUser(id)
		if err == nil {
			c.IndentedJSON(http.StatusOK, user)
		} else {
			c.String(http.StatusNoContent, "")
		}
	} else {
		c.IndentedJSON(http.StatusOK, users)
	}
}

func Create(c *gin.Context) {
	var newUser User

	err := c.BindJSON(&newUser)

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		newId := getNewId(users[len(users)-1].Id)
		if newId != "" {
			newUser.Id = newId
			users = append(users, newUser)
			c.IndentedJSON(http.StatusCreated,
				fmt.Sprintf("User successfally created with id=%s", newId))
		} else {
			c.String(http.StatusBadRequest, "Cannot create a new user")
		}
	}
}

func Update(c *gin.Context) {
	for i, item := range users {
		if item.Id == c.Param("id") {
			c.BindJSON(&item)
			over := users[i+1:]
			users = append(users[:i], item)
			users = append(users, over...)
			c.IndentedJSON(http.StatusOK, item)
			return
		}
	}
	c.String(http.StatusBadRequest, "User not found")
}

func Delete(c *gin.Context) {
	for i, item := range users {
		if item.Id == c.Param("id") {
			users = append(users[:i], users[i+1:]...)
			c.String(http.StatusNoContent, "")
		}
	}
}
