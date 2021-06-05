package main

import (
	"github.com/he2121/go-web/hehe"
	"log"
	"net/http"
)

func main() {
	engine := hehe.New()
	engine.GET("/", indexHandler)
	engine.GET("/hello", helloHandler)
	engine.POST("/login", func(c *hehe.Context) {
		c.JSON(http.StatusOK, hehe.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	log.Fatal(http.ListenAndServe("localhost:1234", engine))
}

func indexHandler(c *hehe.Context) {
	c.String(http.StatusOK, "URL.Path = %s \n", c.Req.URL)
}

func helloHandler(c *hehe.Context)  {
	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
}



