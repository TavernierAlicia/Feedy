package main

import (
	"net/http"
	"strings"
	_"os"
	_"io"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"fmt"
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
	c.Request.ParseForm()
	mail := strings.Join(c.Request.PostForm["mail"], " ")
	name := strings.Join(c.Request.PostForm["name"], " ")
	message := strings.Join(c.Request.PostForm["message"], " ")


	// choose subject and send mail
	// successReq := insertDb(mail, name, message)
	successMail := SendMail(mail, name, message)
	fmt.Println(successMail)
	// if successMail == nil && successReq == nil {
	// 	c.Redirect(http.StatusMovedPermanently, "/index")
	// } else {
	// 	c.Redirect(http.StatusMovedPermanently, "/error")
	// }
	c.Redirect(http.StatusMovedPermanently, "/index")
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

	//launch
	//router.Run(":3000")
	//router.Run()
	//router.Run("www.orderndrink.com:3000")
	router.Run("127.0.0.1:3000")
}
