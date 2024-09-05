package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type User struct {
	ID          uint   `json:"id"`
	Login       string `json:"login"`
	Password    string `json:"password"`
	TokensCount uint64 `json:"tokens_count"`
}

type Message struct {
	SenderId   uint   `json:"sender_id"`
	ReceiverId uint   `json:"receiver_id"`
	Text       string `json:"text"`
}

var users = []User{
	{0, "Some Body", "C00l_p@ss", 200},
	{1, "No Body", "not_so_cool_pass", 20},
	{2, "Just Mind", "I-H@t3_Pa5Se5", 1_000_000},
}

var messages = []Message{
	{0, 1, "Ur password suck"},
	{1, 0, "U haven't seen JustMinds password"},
	{2, 0, "U haven't seen my password"},
}

func getUser(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func getMessages(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, messages)
}

func postPayoutTokens(c *gin.Context) {
	payout, _ := c.GetPostForm("payout")
	payoutUnit, _ := strconv.ParseUint(payout, 10, 64)
	users[0].TokensCount += payoutUnit
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":         "Tapper first page",
		"TokensInPurse": users[0].TokensCount,
	})
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("web/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":         "Tapper first page",
			"TokensInPurse": users[0].TokensCount,
		})
	})
	router.GET("/users", getUser)
	router.GET("/msg", getMessages)
	router.POST("/", postPayoutTokens)
	router.Run("localhost:8081")
}
