package handler

import (
	"encoding/json"
	users "go-app-users"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getUsers(c *gin.Context) {
	jsonFile, err := os.Open("./users.json")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not Found",
		})
		return
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users []users.User

	json.Unmarshal(byteValue, &users)

	c.JSON(http.StatusOK, users)
}
