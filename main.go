package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-crud/controllers"
	"github.com/go-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	initializers.ConnectToDB()
	r := gin.Default() //route

	// // r.GET("/", controllers.PostsCreate)
	// //changing it to POST method, we are creating the data
	r.POST("/post", controllers.PostsCreate)

	// // r.GET("/", func(c *gin.Context) {
	// // 	c.JSON(200, gin.H{
	// // 		"message": "pong",
	// // 	})
	// // })

	r.GET("/post", controllers.PostsRead)

	r.GET("/post/:id", controllers.PostsReadSingle)

	r.PUT("/post/:id", controllers.PostsUpdate)

	r.DELETE("/post/:id", controllers.PostsDelete)

	//Run(":3006")

	r.Run() // listen and serve on 0.0.0.0:8080s
	// controllers.PostsCreate()
	// println("count of entered data is:", result.RowsAffected)

}
