package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddPet - Add a new pet to the store
func AddPet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// DeletePet - Deletes a pet
func DeletePet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// FindPetsByStatus - Finds Pets by status
func FindPetsByStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// FindPetsByTags - Finds Pets by tags
// Deprecated
func FindPetsByTags(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetPetById - Find pet by ID
func GetPetById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
