package main

import (
	gomysql "main/go-mysql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", rootHandler)
	r.GET("/user", userHandler)
	r.GET("/users/:id", userIDHandler)
	r.GET("/query", userCreate)

	r.Run()
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "Hello World!",
	})
}

func userHandler(c *gin.Context) {
	data := gomysql.Testing()
	c.JSON(http.StatusOK, gin.H{
		"user": data,
	})
}

func userIDHandler(c *gin.Context) {
	id := c.Param("id")
	data := gomysql.Testing2(id)
	c.JSON(http.StatusOK, gin.H{
		"user": data,
	})
}

func userCreate(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")
	age := c.Query("age")
	grade := c.Query("grade")

	convAge, _ := strconv.Atoi(age)
	convGrade, _ := strconv.Atoi(grade)

	data := gomysql.Testing3(id, name, convAge, convGrade)
	c.JSON(http.StatusOK, gin.H{
		"user": data,
	})
}
