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
	}

	c.JSON(http.StatusOK, res)
}
