package main

import (
	"github.com/he2121/go-web/hehe"
	"net/http"
)

func main() {
	engine := hehe.New()
	engine.GET("/", indexHandler)
	engine.Use(hehe.Logger())
	engine.Use(hehe.Recovery())
	v1 := engine.Group("/v1")
	v1.GET("/hello", helloHandler)
	v1.GET("/hello/:name", func(c *hehe.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})
	v1.GET("/panic", func(c *hehe.Context) {
		names := []string{"test"}
		c.String(http.StatusOK, names[100])
	})
	v1 = engine.Group("/v2")
	v1.POST("/login", func(c *hehe.Context) {
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

func helloHandler(c *hehe.Context) {
	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
}
