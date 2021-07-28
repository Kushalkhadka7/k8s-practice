package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"os"
	"io/ioutil"
)

// Router is a app level router.
type Router struct {
	*gin.Engine
}

// New initializes new gin router.
func New() *Router {
	gin.ForceConsoleColor()
	r := gin.New()

	r.Use(gin.Recovery())

	r.GET("/info", func(c *gin.Context) {
		fmt.Println(c.Request.Host)

		password,err := ioutil.ReadFile("/data/PASSWORD")
		if err!=nil{
			fmt.Println(err)
		}
		
		userName,err := ioutil.ReadFile("/data/USER_NAME")
		if err!=nil{
			fmt.Println(err)
		}

		c.JSON(200, gin.H{
			"host":    c.Request.Host,
			"url":     c.Request.RequestURI,
			"message": "Success",
			"someEnvs": os.Getenv("NEWNAME"),
			"USER_NAME": string(userName),
			"PASSWORD": string(password),
		})
	})

	r.GET("/login", func(c *gin.Context) {
		fmt.Println(c.Request.Host)
		c.JSON(200, gin.H{
			"host":    c.Request.Host,
			"url":     c.Request.RequestURI,
			"message": "Success",
		})
	})

	return &Router{r}
}
