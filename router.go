package main

import (
	_ "fmt"
	_ "io"
	"net/http"
	_ "os"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//define variables
var (
	log *zap.Logger
)

//handle index page
func index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

//handle form pro
func contact(c *gin.Context) {
	c.HTML(200, "contact.html", nil)
}

func legal(c *gin.Context) {
	c.HTML(200, "legal.html", nil)
}

func receptForm(c *gin.Context) {
	direction := "IN"
	c.Request.ParseForm()
	mail := strings.Join(c.Request.PostForm["mail"], " ")
	name := strings.Join(c.Request.PostForm["name"], " ")
	message := strings.Join(c.Request.PostForm["message"], " ")

	// choose subject and send mail
	successReq := insertDb(mail, name, message, direction)
	successMail := recvMail(mail, name, message)

	if successMail == nil && successReq == nil {
		c.Redirect(http.StatusMovedPermanently, "/index")
	} else {
		c.Redirect(http.StatusMovedPermanently, "/legal")
	}
}

func subscription(c *gin.Context) {
	direction := "OUT"
	c.Request.ParseForm()
	mail := strings.Join(c.Request.PostForm["mail"], " ")
	name := ""
	message := ""

	// choose subject and send mail
	if mail == "" {
		c.Redirect(http.StatusMovedPermanently, "/legal")
	}

	successReq := insertDb(mail, name, message, direction)
	successMail := sendMail(mail, name, message)

	if successMail == nil && successReq == nil {
		c.Redirect(http.StatusMovedPermanently, "/index")
	} else {
		c.Redirect(http.StatusMovedPermanently, "/legal")
	}
}

func main() {

	//zap stuff
	log, _ = zap.NewProduction()
	defer log.Sync()

	//Define router
	router := gin.Default()

	//to include html
	router.LoadHTMLFiles("index.html", "legal.html", "contact.html")

	//to include assets
	router.Static("/assets", "./assets")

	//GET requests
	//index routes
	router.GET("/", index)
	router.GET("/index", index)

	//from infos pages
	router.GET("/contact", contact)
	router.GET("/legal", legal)

	//POST requests
	router.POST("/contact", receptForm)
	router.POST("/index", subscription)
	router.POST("/", subscription)

	//launch
	router.Run(":1234")
}
