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

	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{
			Code:    "INTERNAL_SERVER_ERR",
			Message: "failed to list user posts",
		})
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
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{
			Code:    "BAD_BODY_PARAMS",
			Message: "failed to bind json to structure",
		})
		return
	}

	res, err := h.service.Post().CreatePost(context.Background(), payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{
			Code:    "INTERNAL_SERVER_ERR",
			Message: "failed to list user posts",
		})
		return
	}

	response.ID = res
	c.JSON(http.StatusOK, response)
}
