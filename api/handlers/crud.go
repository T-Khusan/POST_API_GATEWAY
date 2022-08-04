package handlers

import (
	"context"
	"net/http"
	"post_api_gateway/genproto/crud_service"
	"strconv"

	// "post_api_gateway/genproto/post_service"
	// "post_api_gateway/pkg/helper"

	// "contact_api_gateway/genproto/contact_service"

	"github.com/gin-gonic/gin"
)

// GetAllPosts godoc
// @ID get_all_post
// @Router /api/post  [GET]
// @Summary Get all post
// @Description Get All Post
// @Tags Crud
// @Accept json
// @Produce json
// @Param page query integer false "page"
// @Param limit query integer false "limit"
// @Success 200 {object} http.Response{data=crud_service.ListRespPost} "ListRespPostBody"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *Handler) GetAllPost(c *gin.Context) {
	offset, err := h.getOffsetParam(c)
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while getting offset", err.Error())
	}

	limit, err := h.getLimitParam(c)
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while getting limit", err.Error())
	}

	resp, err := h.services.CrudService().GetAll(
		context.Background(),
		&crud_service.ListReq{
			Limit: int64(limit),
			Page:  int64(offset),
		},
	)

	if !handleError(h.log, c, err, "error while getting all posts") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// GetPost godoc
// @ID get_post
// @Router /api/post/{post_id} [GET]
// @Summary get post
// @Description Get Post
// @Tags Crud
// @Accept json
// @Produce json
// @Param post_id path string true "post_id"
// @Success 200 {object} http.Response{data=crud_service.Post} "GetPost"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Internal Server Error"
func (h *Handler) GetPost(c *gin.Context) {
	id := c.Param("post_id")

	post_id, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "parse int error", err.Error())
	}

	resp, err := h.services.CrudService().Get(
		context.Background(),
		&crud_service.ById{
			Id: post_id,
		},
	)

	if !handleError(h.log, c, err, "error while getting post") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// UpdatePost godoc
// @ID update_post
// @Router /api/post/{post_id} [PUT]
// @Summary Update Post
// @Description Update Post by ID
// @Tags Crud
// @Accept json
// @Produce json
// @Param post_id path string true "post_id"
// @Param post body crud_service.UpdatePost true "UpdatePostBody"
// @Success 200 {object} http.Response{data=crud_service.Post} "Post data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdatePost(c *gin.Context) {
	var post crud_service.UpdatePost

	post_id := c.Param("post_id")
	p_id, err := strconv.ParseInt(post_id, 10, 32)
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "parse int error", err.Error())
	}

	if err := c.BindJSON(&post); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.CrudService().Update(
		context.Background(),
		&crud_service.Post{
			Id:     p_id,
			UserId: post.UserId,
			Title:  post.Title,
			Body:   post.Body,
		},
	)

	if !handleError(h.log, c, err, "error while getting post") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// DeletePost godoc
// @ID delete_post
// @Router /api/post/{post_id} [DELETE]
// @Summary Delete Post
// @Description Delete Post by given ID
// @Tags Crud
// @Accept json
// @Produce json
// @Param post_id path string true "post_id"
// @Success 200 {object} http.Response{data=string} "Post deleted"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeletePost(c *gin.Context) {
	post_id := c.Param("post_id")

	id, err := strconv.ParseInt(post_id, 10, 32)
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "parse int error", err.Error())
	}

	resp, err := h.services.CrudService().Delete(
		context.Background(),
		&crud_service.ById{
			Id: id,
		},
	)
	if !handleError(h.log, c, err, "error while deleting post") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}
