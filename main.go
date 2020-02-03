package main

import "github.com/gin-gonic/gin"

interface Todolist {

}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(ctx *gin.Context) {
		todolist := {
			name: "todoname",
			detail: "detail_text"
		}
		ctx.JSON(200, todolist)
	})

	router.Run()
}
