package controllers

import (
	"mini-project/models"
	"mini-project/repositories"
	services "mini-project/services/blogs"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BlogController struct {
	Service services.BlogService
}

func NewBlogController() BlogController {
	return BlogController{
		Service: services.NewBlogService(&repositories.BlogRepositoryImpl{}),
	}
}

func (bc *BlogController) GetAll(c echo.Context) error {
	var blogs []models.Blog
	blogs = bc.Service.Repository.GetAll()

	var blogsResponse []models.BlogResponse

	for _, blog := range blogs {
		blogsResponse = append(blogsResponse, blog.ToResponse())
	}

	return NewResponseSuccess(c, http.StatusOK, "successfully get all blogs", blogsResponse)
}

func (bc *BlogController) GetByUserID(c echo.Context) error {
	userIDString := c.Param("user_id")
	userID, _ := strconv.Atoi(userIDString)

	var blogs []models.Blog
	blogs = bc.Service.Repository.GetByUserID(userID)

	var blogsResponse []models.BlogResponse

	for _, blog := range blogs {
		blogsResponse = append(blogsResponse, blog.ToResponse())
	}

	return NewResponseSuccess(c, http.StatusOK, "successfully get blogs by user id", blogsResponse)
}

func (bc *BlogController) GetByCategoryID(c echo.Context) error {
	categoryIDString := c.Param("category_id")
	categoryID, _ := strconv.Atoi(categoryIDString)

	var blogs []models.Blog
	blogs = bc.Service.Repository.GetByCategoryID(categoryID)

	var blogsResponse []models.BlogResponse

	for _, blog := range blogs {
		blogsResponse = append(blogsResponse, blog.ToResponse())
	}

	return NewResponseSuccess(c, http.StatusOK, "successfully get blogs by category id", blogsResponse)
}

func (bc *BlogController) GetByID(c echo.Context) error {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	var blog models.Blog

	blog = bc.Service.Repository.GetByID(id)

	if blog.ID == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "blog doesn't exist")
	}

	return NewResponseSuccess(c, http.StatusOK, "successfully get blog", blog.ToResponse())
}

func (bc *BlogController) Create(c echo.Context) error {
	var blogRequest models.BlogRequest

	c.Bind(&blogRequest)

	if err := blogRequest.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "validation failed")
	}

	blog := bc.Service.Repository.Create(blogRequest)

	return NewResponseSuccess(c, http.StatusOK, "successfully create blog", blog.ToResponse())
}

func (bc *BlogController) Update(c echo.Context) error {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	var blogUpdate models.BlogRequest

	c.Bind(&blogUpdate)

	if err := blogUpdate.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "validation failed")
	}

	blog := bc.Service.Repository.Update(id, blogUpdate)

	if blog.ID == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "blog doesn't exist")
	}

	return NewResponseSuccess(c, http.StatusOK, "successfully update blog", blog.ToResponse())
}

func (bc *BlogController) Delete(c echo.Context) error {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	isDeleted := bc.Service.Repository.Delete(id)

	if !isDeleted {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "unsuccessfully delete blog",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "successfully delete blog",
	})
}
