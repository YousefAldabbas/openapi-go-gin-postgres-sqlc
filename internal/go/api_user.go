/*
 * OpenAPI Petstore
 *
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUser - Create user
func CreateUser(c *gin.Context) {
	var user User
	c.BindJSON(&user)
	c.String(http.StatusOK, "Would have created a user for %v", user.Username)
}

// CreateUsersWithArrayInput - Creates list of users with given input array
func CreateUsersWithArrayInput(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// CreateUsersWithListInput - Creates list of users with given input array
func CreateUsersWithListInput(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteUser - Delete user
func DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetUserByName - Get user by user name
func GetUserByName(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// LoginUser - Logs user into the system
func LoginUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// LogoutUser - Logs out current logged in user session
func LogoutUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// UpdateUser - Updated user
func UpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
