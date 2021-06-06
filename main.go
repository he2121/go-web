package main

import (
	"github.com/he2121/go-web/hehe"
	"net/http"
)

func main() {
	engine := hehe.New()
	engine.GET("/", indexHandler)
	engine.GET("/hello", helloHandler)
	engine.GET("hello/:name", func(c *hehe.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})
	engine.POST("/login", func(c *hehe.Context) {
		c.JSON(http.StatusOK, hehe.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	engine.Run(":1234")
}

func indexHandler(c *hehe.Context) {
	c.String(http.StatusOK, "URL.Path = %s \n", c.Req.URL)
}

func helloHandler(c *hehe.Context)  {
	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
}



