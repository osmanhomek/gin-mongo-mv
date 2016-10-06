package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

type (
	// SectionControllerImpl represents the controller for operating on the User resource
	SectionControllerImpl struct {
		session *mgo.Session
	}
)

// NewSectionController provides a reference to a UserController with provided mongo session
func NewSectionController(s *mgo.Session) *SectionControllerImpl {
	return &SectionControllerImpl{s}
}

// GetSections retrieves an individual user resource
func (uc SectionControllerImpl) GetSections(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
}
