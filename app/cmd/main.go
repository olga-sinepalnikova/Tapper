package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"tapper/internal/storage"
)

func getUser(c *gin.Context) {
	//c.IndentedJSON(http.StatusOK, users)
}

func getMessages(c *gin.Context) {
	//c.IndentedJSON(http.StatusOK, messages)
}

func postPayoutTokens(c *gin.Context) {
	//payout, _ := c.GetPostForm("payout")
	//payoutUnit, _ := strconv.ParseUint(payout, 10, 64)
	//users[0].TokensCount += payoutUnit
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":         "Tapper first page",
		"TokensInPurse": 0,
	})
}

func main() {
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)

	db, err := storage.Connect()
	if err != nil {
		log.Fatal(err)
		os.Exit(69)
	}
	log.Info("Connected to database")
	log.Debugln(db)

	storage.GenerateStructs(db, log)

	router := gin.Default()
	router.LoadHTMLGlob("web/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":         "Tapper first page",
			"TokensInPurse": 0,
		})
	})
	router.GET("/users", getUser)
	router.GET("/msg", getMessages)
	router.POST("/", postPayoutTokens)

	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
	log.Debug("Server started")
	log.Debug("Listening on port 8080")
}
