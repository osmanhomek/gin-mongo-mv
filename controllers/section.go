package controllers

import (
	"net/http"
	//"encoding/json"
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"gin-mongo-mv/models"
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

type Sections struct {  
    orderIndex    int   `json:"orderIndex"`
    engName   string   `json:"engName"`
    trName string `json:"trName"`
    avatar   string   `json:"avatar"`
}

// GetSections retrieves an individual user resource
func (uc SectionControllerImpl) GetSections(c *gin.Context) {

	u := make([]models.Section,0,3)
	//result := make([]models.User, 0, 10) 
	//err := uc.session.DB("gunde5dkcom").C("sections").Find(bson.M{"engName": "fruits"}).Sort("orderIndex").All(&u)
	err := uc.session.DB("gunde5dkcom").C("sections").Find(nil).Sort("orderIndex").All(&u)
	if err != nil {
		log.Println("Failed get all sections: ", err)
        return
	}
	fmt.Println(u)

	/*
	sessionc := uc.session.Copy()
	defer sessionc.Close()
	c := sessionc.DB("gunde5dkcom").C("sections")

	var sections []Sections
	err := c.Find(bson.M{}).All(&sections)  
	if err != nil {
		log.Println("Failed get all books: ", err)
        return
	}

	respBody, err := json.MarshalIndent(sections, "", "  ")
	if err != nil {
    	log.Println("Failed get all books: ", err)
	}	
	fmt.Println(respBody)
	*/
	fmt.Println("fmt")
	log.Println("log")
	//c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	c.JSON(http.StatusOK, u)
}
