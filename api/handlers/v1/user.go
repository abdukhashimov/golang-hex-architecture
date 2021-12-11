package v1

import (
	"context"
	"net/http"

	"github.com/abdukhashimov/golang-hex-architecture/service/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func (h *handlerV1) GetUserPosts(c *gin.Context) {
	author := c.DefaultQuery("author", "")
	offset := cast.ToInt(c.DefaultQuery("offset", "0"))
	limit := cast.ToInt(c.DefaultQuery("limit", "10"))

	res, err := h.service.Post().GetPostsByAuthhor(context.Background(), models.PostFilter{
		Author: author,
		Offset: offset,
		Limit:  limit,
	})

	if h.handleInternal(c, "failed to get posts by author", err) {
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *handlerV1) CreateUserPost(c *gin.Context) {
	var (
		payload  models.Post
		response ResponseCreated
	)

	err := c.ShouldBindJSON(&payload)
	if h.handleBadRequest(c, "failed to convert json to struct", err) {
		return
	}

	res, err := h.service.Post().CreatePost(context.Background(), payload)
	if h.handleInternal(c, "failed to create post", err) {
		return
	}

	response.ID = res
	c.JSON(http.StatusOK, response)
}
