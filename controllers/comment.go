package controllers

import (
	"mini-project/models"
	"mini-project/repositories"
	services "mini-project/services/comments"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CommentController struct {
	Service services.CommentService
}

func NewCommentController() CommentController {
	return CommentController{
		Service: services.NewCommentService(&repositories.CommentRepositoryImpl{}),
	}
}

func (cc *CommentController) GetAll(c echo.Context) error {
	var comments []models.Comment
	comments = cc.Service.Repository.GetAll()

	var commentsResponse []models.CommentResponse

	for _, comment := range comments {
		commentsResponse = append(commentsResponse, comment.ToResponse())
	}

	return NewResponseSuccess(c, http.StatusOK, "successfully get all comments", commentsResponse)
}

func (cc *CommentController) GetByBlogID(c echo.Context) error {
	blogIDString := c.Param("blog_id")

	blogID, _ := strconv.Atoi(blogIDString)

	var comments []models.Comment

	comments = cc.Service.Repository.GetByBlogID(blogID)

	var commentsResponse []models.CommentResponse

	for _, comment := range comments {
		commentsResponse = append(commentsResponse, comment.ToResponse())
	}

	return NewResponseSuccess(c, http.StatusOK, "successfully get comments by blog id", commentsResponse)
}

func (cc *CommentController) GetByID(c echo.Context) error {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	var comment models.Comment

	comment = cc.Service.Repository.GetByID(id)

	if comment.ID == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "comment doesn't exist")
	}

	return NewResponseSuccess(c, http.StatusOK, "successfully get comment", comment.ToResponse())
}

func (cc *CommentController) Create(c echo.Context) error {
	var commentRequest models.CommentRequest

	c.Bind(&commentRequest)

	if err := commentRequest.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "validation failed")
	}

	comment := cc.Service.Repository.Create(commentRequest)

	return NewResponseSuccess(c, http.StatusOK, "successfully create comment", comment.ToResponse())
}

func (cc *CommentController) Update(c echo.Context) error {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	var commentUpdate models.CommentRequest

	c.Bind(&commentUpdate)

	if err := commentUpdate.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "validation failed")
	}

	comment := cc.Service.Repository.Update(id, commentUpdate)

	if comment.ID == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "comment doesn't exist")
	}

	return NewResponseSuccess(c, http.StatusOK, "successfully update comment", comment.ToResponse())
}

func (cc *CommentController) Delete(c echo.Context) error {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	isDeleted := cc.Service.Repository.Delete(id)

	if !isDeleted {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "unsuccessfully delete comment",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "successfully delete comment",
	})
}
