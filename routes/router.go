package routes

import (
	"net/http"
	"strconv"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"

	"untitledgoproject/models"
	"untitledgoproject/controllers"
)

func getAllUsers(db *gorm.DB) []models.User {
	var users []models.User
	db.Find(&users)
	return users
}

func render(c *gin.Context, template string) {
	c.HTML(http.StatusOK, "base.html", gin.H{"content": template})
}

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	userController := controllers.NewUserController(db)

	// Load HTML templates
	r.LoadHTMLGlob("templates/*.html")
	// Serve static files from the "public" directory
  r.Static("/public", "public")

	r.GET("/", func(c *gin.Context) {
		render(c, "index.html")
	})

	r.GET("/api/users", func(c *gin.Context) {
		users := userController.GetAllUsers()
		c.JSON(http.StatusOK, users)
	})

	r.POST("/api/users", func(c *gin.Context) {
		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := userController.CreateUser(&user); err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, user)
	})
	
	r.PUT("/api/users/:id", func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := userController.EditUser(uint(id), &user); err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)
	})

	r.DELETE("/api/users/:id", func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := userController.DeleteUser(uint(id)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusNoContent, nil)
	})

	return r
}