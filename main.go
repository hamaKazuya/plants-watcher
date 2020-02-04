package main

// import "github.com/jinzhu/gorm"
import "github.com/gin-gonic/gin"

type User struct {
	// gorm.Model
	Name   string `json:"name"`
	Detail string `json:"detail"`
}

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		User := User{}
		User.Name = "Username"
		User.Detail = "detail_text"
		ctx.JSON(200, User)
	})

	router.GET("/plants", func(ctx *gin.Context) {
		User := User{"Username", "detail_text"}
		ctx.JSON(200, User)
	})

	router.Run(":8088")
}
