package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seigaalghi/firestore-go/models"
	"github.com/seigaalghi/firestore-go/repository"
)

// GetPosts is to get All Posts from database
func GetPosts(ctx *gin.Context) {
	data, err := repository.FindAllPosts()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// AddPost is to create a new post
func AddPost(ctx *gin.Context) {
	var input models.Post
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	data, err := repository.SavePost(&models.Post{
		Title:     input.Title,
		Text:      input.Text,
		Date:      input.Date,
		Price:     input.Price,
		Authors:   input.Authors,
		Published: input.Published,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// EditPost is to edit a post
func EditPost(ctx *gin.Context) {
	var input models.Post
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	data, err := repository.UpdatePost(&models.Post{
		Title:     input.Title,
		Text:      input.Text,
		Date:      input.Date,
		Price:     input.Price,
		Authors:   input.Authors,
		Published: input.Published,
	}, ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// DeletePost is to delet a post
func DeletePost(ctx *gin.Context) {
	err := repository.RemovePost(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Post deleted successfully",
	})
}
