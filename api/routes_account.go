package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/RTradeLtd/Temporal/models"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var nilTime time.Time

func ChangeAccountPassword(c *gin.Context) {
	ethAddress := GetAuthenticatedUserFromContext(c)

	oldPassword, exists := c.GetPostForm("old_password")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "old_password post form param does not exist",
		})
		return
	}

	newPassword, exists := c.GetPostForm("new_password")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "new_password post form param does not exist",
		})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	um := models.NewUserManager(db)

	suceeded, err := um.ChangePassword(ethAddress, oldPassword, newPassword)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("error occured while changing password %s", err.Error()),
		})
		return
	}
	if !suceeded {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "password change failed but no error occured",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "password changed",
	})
}

// RegisterUserAccount is used to sign up with temporal and gain web interface access
// you will not be granted API access however, as that needs to be done manually
func RegisterUserAccount(c *gin.Context) {
	ethAddress, exists := c.GetPostForm("eth_address")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "eth_address parameter does not exist"})
		return
	}
	password, exists := c.GetPostForm("password")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password parameter does not exist"})
		return
	}
	db := c.MustGet("db").(*gorm.DB)

	userManager := models.NewUserManager(db)
	userModel, err := userManager.NewUserAccount(ethAddress, password, false)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userModel.HashedPassword = "scrubbed"
	c.JSON(http.StatusCreated, gin.H{"user": userModel})
	return
}

// RegisterEnterpriseUserAccount is used to register a user account marked as enterprise enabled
func RegisterEnterpriseUserAccount(c *gin.Context) {
	ethAddress, exists := c.GetPostForm("eth_address")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "eth_address parameter does not exist"})
		return
	}
	password, exists := c.GetPostForm("password")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password parameter does not exist"})
		return
	}
	db := c.MustGet("db").(*gorm.DB)

	userManager := models.NewUserManager(db)
	userModel, err := userManager.NewUserAccount(ethAddress, password, false)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userModel.HashedPassword = "scrubbed"
	c.JSON(http.StatusCreated, gin.H{"user": userModel})
	return
}

// GetAuthenticatedUserFromContext is used to pull the eth address of hte user
func GetAuthenticatedUserFromContext(c *gin.Context) string {
	claims := jwt.ExtractClaims(c)
	// this is their eth address
	return claims["id"].(string)
}
