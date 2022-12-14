package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Blog struct {
	ID         int            `json:"id"`
	UserID     int            `json:"user_id"`
	User       User           `json:"user"`
	CategoryID int            `json:"category_id"`
	Category   Category       `json:"category"`
	Title      string         `json:"title"`
	Content    string         `json:"content"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
	Comments   []Comment      `json:"comment" gorm:"constraint:OnDelete:CASCADE;"`
	Likes      []Like         `json:"likes" gorm:"constraint:OnDelete:CASCADE;"`
}

func (b *Blog) ToResponse() BlogResponse {
	var commentsResponse []CommentResponse
	var likesResponse []LikeResponse

	for _, comment := range b.Comments {
		commentsResponse = append(commentsResponse, comment.ToResponse())
	}

	for _, like := range b.Likes {
		likesResponse = append(likesResponse, like.ToResponse())
	}

	return BlogResponse{
		ID:           b.ID,
		UserID:       b.UserID,
		FirstName:    b.User.FirstName,
		LastName:     b.User.LastName,
		CategoryName: b.Category.Name,
		Title:        b.Title,
		Content:      b.Content,
		Comments:     commentsResponse,
		CreatedAt:    b.CreatedAt,
		UpdatedAt:    b.UpdatedAt,
		LikeAmount:   len(likesResponse),
		Likes:        likesResponse,
	}
}

type BlogRequest struct {
	UserID     int    `json:"user_id" validate:"required"`
	CategoryID int    `json:"category_id" validate:"required"`
	Title      string `json:"title" validate:"required"`
	Content    string `json:"content" validate:"required"`
}

func (br *BlogRequest) ToDBForm() Blog {
	return Blog{
		UserID:     br.UserID,
		CategoryID: br.CategoryID,
		Title:      br.Title,
		Content:    br.Content,
	}
}

func (br *BlogRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(br)

	return err
}

type BlogResponse struct {
	ID           int               `json:"id"`
	UserID       int               `json:"user_id"`
	FirstName    string            `json:"first_name"`
	LastName     string            `json:"last_name"`
	CategoryName string            `json:"category_name"`
	Title        string            `json:"title"`
	Content      string            `json:"content"`
	CreatedAt    time.Time         `json:"created_at"`
	UpdatedAt    time.Time         `json:"updated_at"`
	DeletedAt    time.Time         `json:"deleted_at"`
	Comments     []CommentResponse `json:"comments"`
	LikeAmount   int               `json:"like_amount"`
	Likes        []LikeResponse    `json:"likes"`
}
