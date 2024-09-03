package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	ID          uint   `json:"id"`
	Login       string `json:"login"`
	Password    string `json:"password"`
	TokensCount uint   `json:"tokens_count"`
}

var users = []User{
	{0, "Some Body", "C00l_p@ss", 200},
	{1, "No Body", "not_so_cool_pass", 20},
	{2, "Just Mind", "I-H@t3_Pa5Se5", 1_000_000},
}

func getUser(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func main() {
	router := gin.Default()
	router.GET("/", getUser)
	router.Run("localhost:8081")
}
