package main

import (
	_ "fmt"
	_ "io"
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
	c.HTML(200, "index.html", map[string]interface{}{"send": 0, "ok": 0})
}

//handle form pro
func contact(c *gin.Context) {
	c.HTML(200, "contact.html", map[string]interface{}{"send": 0, "ok": 0})
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
		c.HTML(200, "contact.html", map[string]interface{}{"send": 1, "ok": 1})
	} else {
		c.HTML(200, "contact.html", map[string]interface{}{"send": 0, "ok": 1})
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
		c.HTML(200, "index.html", map[string]interface{}{"send": 0, "ok": 1})
		return
	}

	successReq := insertDb(mail, name, message, direction)
	successMail := sendMail(mail, name, message)

	if successMail == nil && successReq == nil {
		c.HTML(200, "index.html", map[string]interface{}{"send": 1, "ok": 1})
	} else {
		c.HTML(200, "index.html", map[string]interface{}{"send": 0, "ok": 1})
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
