package main

import "github.com/gin-gonic/gin"

type Todo struct {
	Name   string `json:"name"`
	Detail string `json:"detail"`
}

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		todo := Todo{}
		todo.Name = "todoname"
		todo.Detail = "detail_text"
		ctx.JSON(200, todo)
	})

	router.GET("/plants", func(ctx *gin.Context) {
		todo := Todo{"todoname", "detail_text"}
		ctx.JSON(200, todo)
	})

	router.Run(":8080")
}
