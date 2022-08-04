package handlers

import (
	"context"
	"fmt"
	"net/http"
	"post_api_gateway/genproto/post_service"
	"post_api_gateway/pkg/helper"

	"github.com/gin-gonic/gin"
)

// ReloadPost godoc
// @Summary Reload post
// ID reload_post
// @Tags Post
// @Accept json
// @Produce json
// @Success 200 {object} http.Response{data=string} "Post data"
// @Failure 400 {object} http.Response{data=string} "Bad request"
// @Failure 500 {object} http.Response{data=string} "Internal server error"
// @Router /api/post/reload [PUT]
func (h *Handler) ReloadPost(c *gin.Context) {
	var postsList []*post_service.Post

	for i := 1; i < 170; i++ {
		var posts post_service.PostReloadReq

		urlbody := fmt.Sprintf("%v/public/v1/posts/?page=%v", h.cfg.Endpoint, i)
		code, resp, err := helper.MakeRequests(http.MethodGet, urlbody, nil)

		if err != nil {
			h.handleErrorResponse(c, code, "error while loading endpoint json", err)
			return
		}

		fmt.Println("RESP ---> ", resp)
		err = helper.JsonToJson(&posts, resp)
		if err != nil {
			h.handleErrorResponse(c, http.StatusBadRequest, "error while jsontojson", err.Error())
			return
		}

		postsList = append(postsList, posts.Data...)
	}

	_, err := h.services.PostService().Reload(
		context.Background(),
		&post_service.PostReloadReq{
			Data: postsList,
		},
	)

	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while reloading service", err.Error())
		return
	}
	h.handleSuccessResponse(c, http.StatusOK, "ok", postsList)
}
