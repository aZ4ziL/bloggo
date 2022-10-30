package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/aZ4ziL/bloggo/models"
	"github.com/aZ4ziL/bloggo/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type APIV1 interface {
	GetAllCategories() gin.HandlerFunc

	// Article
	Article() gin.HandlerFunc

	// Comment
	GetAllCommentsByArticleID() gin.HandlerFunc

	// User
	Register() gin.HandlerFunc
	GetUserByUsername() gin.HandlerFunc
	Login() gin.HandlerFunc
	Logout() gin.HandlerFunc
}

type apiV1 struct{}

func NewAPIV1() APIV1 {
	return &apiV1{}
}

// GetAllArticle Restfull API return all object of articles
func (a apiV1) Article() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)

		user := session.Get("user")
		if user == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  "not_authenticated",
				"message": "You not authenticated",
			})
			return
		}

		if ctx.Request.Method == "GET" {
			if slug, ok := ctx.GetQuery("slug"); ok {
				article, err := models.GetArticleBySlug(slug)
				if err != nil {
					http.Error(ctx.Writer, err.Error(), http.StatusNotFound)
					return
				}
				ctx.JSON(http.StatusOK, article)
				return
			}

			articles := models.GetAllArticles()
			ctx.JSON(http.StatusOK, articles)
		}

		// Method POST to create new article
		if ctx.Request.Method == "POST" {
			categoryIDInt, _ := strconv.Atoi(ctx.PostForm("category_id"))
			authorIDInt, _ := strconv.Atoi(ctx.PostForm("author_id"))

			var filename string
			logo, err := ctx.FormFile("logo")
			if err != nil {
				filename = ""
			} else {
				filename = "media/article/" + logo.Filename
			}

			article := &models.Article{
				CategoryID: uint(categoryIDInt),
				AuthorID:   uint(authorIDInt),
				Title:      ctx.PostForm("title"),
				Slug:       ctx.PostForm("slug"),
				Logo:       "/" + filename,
				Desc:       ctx.PostForm("desc"),
				Content:    ctx.PostForm("content"),
				Status:     ctx.PostForm("status"),
			}

			validate = validator.New()

			err = validate.Struct(article)
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
				ctx.JSON(http.StatusOK, gin.H{
					"status":   "error",
					"messages": errorMessages,
				})
				return
			}

			err = models.CreateNewArticle(article)
			if err != nil {
				http.Error(ctx.Writer, err.Error(), http.StatusBadRequest)
				return
			}

			err = ctx.SaveUploadedFile(logo, filename)
			if err != nil {
				http.Error(ctx.Writer, err.Error(), http.StatusBadRequest)
				return
			}

			ctx.JSON(http.StatusCreated, gin.H{
				"status": "success",
			})
			return
		}

		// Method PUT to update the article
		if ctx.Request.Method == "PUT" {
			if id, ok := ctx.GetQuery("id"); ok {
				idInt, _ := strconv.Atoi(id)

				article, err := models.GetArticleByID(uint(idInt))
				if err != nil {
					http.Error(ctx.Writer, err.Error(), http.StatusNotFound)
					return
				}

				if utils.GetSessionValue(user)["id"].(uint) != article.AuthorID {
					ctx.JSON(http.StatusForbidden, gin.H{
						"status":  "error",
						"message": "Permission danied.",
					})
					return
				}

				if err != nil {
					http.Error(ctx.Writer, err.Error(), http.StatusNotFound)
					return
				}

				logo, err := ctx.FormFile("logo")
				if err == nil {
					// Delete old files
					filename := strings.Replace(article.Logo, "/", "", 1)
					os.Remove(filename)
				}

				filename := "media/article/" + logo.Filename

				article.Title = ctx.PostForm("title")
				article.Slug = ctx.PostForm("slug")
				article.Logo = "/" + filename
				article.Desc = ctx.PostForm("desc")
				article.Content = ctx.PostForm("content")
				article.Status = ctx.PostForm("status")

				models.GetDB().Save(&article)

				if err = ctx.SaveUploadedFile(logo, filename); err != nil {
					http.Error(ctx.Writer, err.Error(), http.StatusBadRequest)
					return
				}

				ctx.JSON(http.StatusCreated, gin.H{
					"status":  "success",
					"message": "Successfully to update this data.",
				})
				return
			}
		}

		// Mehotd DELETE to delete the article
		if ctx.Request.Method == "DELETE" {
			if id, ok := ctx.GetQuery("id"); ok {
				idInt, _ := strconv.Atoi(id)

				article, err := models.GetArticleByID(uint(idInt))
				if err != nil {
					http.Error(ctx.Writer, err.Error(), http.StatusNotFound)
					return
				}

				if utils.GetSessionValue(user)["id"].(uint) != article.AuthorID {
					ctx.JSON(http.StatusForbidden, gin.H{
						"status":  "error",
						"message": "Permission danied.",
					})
					return
				}

				models.GetDB().Delete(&article)
				ctx.JSON(http.StatusOK, gin.H{
					"status":  "success",
					"message": "Successfully to delete the article.",
				})
				return
			}
		}
	}
}

// GetAllComments Restfull api return All comments
func (a apiV1) GetAllCommentsByArticleID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if articleID, ok := ctx.GetQuery("article_id"); ok {
			articleIDInt, _ := strconv.Atoi(articleID)
			comments := models.GetAllCommentsByArticleID(uint(articleIDInt))
			ctx.JSON(http.StatusOK, comments)
			return
		}
	}
}

// GetAllCategories Restfull API return all object of categories
func (a apiV1) GetAllCategories() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if slug, ok := ctx.GetQuery("slug"); ok {
			category, err := models.GetCategoryBySlug(slug)
			if err != nil {
				http.Error(ctx.Writer, err.Error(), http.StatusNotFound)
				return
			}
			ctx.JSON(http.StatusOK, category)
			return
		}
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
				"status":  "success",
				"message": "User with username: " + user.Username + " is registered.",
			})
			return
		} else {
			http.Error(ctx.Writer, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
	}
}

// GetUserByUsername Restfull API with Method GET and Query user
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

// Login Restfull for login user
func (a apiV1) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)

		// If user is authenticated
		if user := session.Get("user"); user != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"status":  "authenticated",
				"message": "You has already authenticated.",
			})
			return
		}

		if ctx.Request.Method == "GET" {
			http.Error(ctx.Writer, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		if ctx.Request.Method == "POST" {
			username := ctx.PostForm("username")
			password := ctx.PostForm("password")

			user, err := models.GetUserByUsername(username)
			if err != nil {
				ctx.JSON(http.StatusOK, gin.H{
					"status":  "error",
					"message": "Username not register.",
				})
				return
			}

			if !utils.DecryptionPassword(user.Password, password) {
				ctx.JSON(http.StatusOK, gin.H{
					"status":  "error",
					"message": "Password is wrong.",
				})
				return
			}

			userSession := map[string]interface{}{
				"id":          user.ID,
				"firstName":   user.FirstName,
				"lastName":    user.LastName,
				"fullName":    user.FirstName + " " + user.LastName,
				"email":       user.Email,
				"username":    user.Username,
				"isSuperuser": user.IsSuperuser,
				"isStaff":     user.IsStaff,
				"isActive":    user.IsActive,
				"dateJoined":  user.DateJoined,
				"lastLogin":   user.LastLogin,
			}

			session.Set("user", userSession)
			if err := session.Save(); err != nil {
				http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
				return
			}

			ctx.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": "you are now authenticated.",
			})
			return
		}
	}
}

// Logout Restfull api for logout user
func (a apiV1) Logout() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)

		session.Delete("user")
		session.Clear()
		if err := session.Save(); err != nil {
			http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "You logged out.",
		})
	}
}
