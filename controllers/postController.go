package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-crud/initializers"
	"github.com/go-crud/models"
)

//
// can be declared outside of PostsCreate
// type User struct {
// 	Title string
// 	Body  string
// }

// **************************************************#1
func PostsCreate(c *gin.Context) {

	//loading data off req body
	var user struct {
		Title string
		Body  string
	}

	//create a post
	//post := models.Post{Title: "hello", Body: "post body"} //hardcoding the data for now
	// var user User

	c.Bind(&user)
	post := models.Post{Title: user.Title, Body: user.Body}

	//crerating in initializers
	result := initializers.DB.Create(&post) // pass pointer of data to Create
	if result.Error != nil {
		c.Status(400)
		fmt.Println("Error Found!!")

	}
	//return it or responding it
	c.JSON(200, gin.H{
		//"post": user,  //this will not return the "id","createdAT", "updatedat".....
		"post": post,
	})
}

// **************************************************#2
func PostsRead(c *gin.Context) {
	//get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	//respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

// **************************************************#3
func PostsReadSingle(c *gin.Context) {
	//get id from url
	id := c.Param("id")

	//get the posts
	var post models.Post
	initializers.DB.First(&post, id)

	//respond with them
	c.JSON(200, gin.H{
		"posts": post,
	})
}

// **************************************************#4
func PostsUpdate(c *gin.Context) {
	//get the id from url
	id := c.Param("id")

	//get the data off req body
	var user struct {
		Title string
		Body  string
	}
	c.Bind(&user)

	//finding the post to update
	var post models.Post
	initializers.DB.First(&post, id)

	//updating
	initializers.DB.Model(&post).Updates(models.Post{
		Title: user.Title,
		Body:  user.Body,
	})

	//repond with it
	c.JSON(200, gin.H{
		"posts": post,
	})
}

// **************************************************#5
func PostsDelete(c *gin.Context) {
	//get id from the url
	id := c.Param("id")

	//delete
	initializers.DB.Delete(&models.Post{}, id)

	//respond with it
	c.Status(200)
}
