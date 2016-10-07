package controllers

import (
	"net/http"
	"log"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
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

	var m []bson.M
	err := uc.session.DB("gunde5dkcom").C("sections").Find(nil).All(&m)
	if err != nil {
		log.Println("Failed get all sections: ", err)
		return
	}	

	for _, value := range m {
		fmt.Println(value["trName"])
	    for jsonKey, jsonValue := range value {
	    	fmt.Println(jsonKey, ":", jsonValue)
	    }
	    fmt.Println("\n")
	}	

	log.Println("log")
	//c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	c.JSON(http.StatusOK, m)
}
