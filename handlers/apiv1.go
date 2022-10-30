package handlers

import (
	"fmt"
	"net/http"

	"github.com/aZ4ziL/bloggo/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type APIV1 interface {
	GetAllCategories() gin.HandlerFunc
	GetAllArticles() gin.HandlerFunc

	// User
	Register() gin.HandlerFunc
	GetUserByUsername() gin.HandlerFunc
}

type apiV1 struct{}

func NewAPIV1() APIV1 {
	return &apiV1{}
}

// GetAllArticle Restfull API return all object of articles
func (a apiV1) GetAllArticles() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		articles := models.GetAllArticles()
		ctx.JSON(http.StatusOK, articles)
	}
}

// GetAllCategories Restfull API return all object of categories
func (a apiV1) GetAllCategories() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		categories := models.GetAllCategories()
		ctx.JSON(http.StatusOK, categories)
	}
}

// Register Restfull API with method POST to registering a user.
func (a apiV1) Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Method == "POST" {
			validate = validator.New()

			user := &models.User{
				FirstName: ctx.PostForm("first_name"),
				LastName:  ctx.PostForm("last_name"),
				Username:  ctx.PostForm("username"),
				Email:     ctx.PostForm("email"),
				Password:  ctx.PostForm("password"),
			}

			err := validate.Struct(user)
			if err != nil {
				if _, ok := err.(*validator.InvalidValidationError); ok {
					fmt.Println(err)
					return
				}

				errorMessages := []string{}
				for _, err := range err.(validator.ValidationErrors) {
					errorMessage := fmt.Sprintf("Error at %s with %s.", err.Field(), err.ActualTag())
					errorMessages = append(errorMessages, errorMessage)
				}

				ctx.JSON(http.StatusBadRequest, gin.H{
					"errors": errorMessages,
				})
				return
			}

			err = models.CreateNewUser(user)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"errors": err.Error(),
				})
				return
			}

			ctx.JSON(http.StatusCreated, gin.H{
				"status": "User with username: " + user.Username + " is registered.",
			})
			return
		} else {
			http.Error(ctx.Writer, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
	}
}

func (a apiV1) GetUserByUsername() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Method == "GET" {
			if username, ok := ctx.GetQuery("username"); ok {
				user, err := models.GetUserByUsername(username)
				if err != nil {
					http.Error(ctx.Writer, err.Error(), http.StatusNotFound)
					return
				}

				ctx.JSON(http.StatusOK, gin.H{
					"id":           user.ID,
					"first_name":   user.FirstName,
					"last_name":    user.LastName,
					"username":     user.Username,
					"email":        user.Email,
					"is_superuser": user.IsSuperuser,
					"is_staff":     user.IsStaff,
					"is_active":    user.IsActive,
					"last_login":   user.LastLogin,
					"date_joined":  user.DateJoined,
					"articles":     user.Articles,
					"comments":     user.Comments,
				})
				return
			}
		}
	}
}
