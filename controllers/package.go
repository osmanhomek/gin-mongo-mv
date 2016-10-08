package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	// PackageControllerImpl represents the controller for operating on the User resource
	PackageControllerImpl struct {
		session *mgo.Session
	}
	// Package asds
	Package struct {
		Sections []map[string]interface{} `json:"sections" bson:"sections"`
	}
)

// NewPackageController provides a reference to a UserController with provided mongo session
func NewPackageController(s *mgo.Session) *PackageControllerImpl {
	return &PackageControllerImpl{s}
}

// GetPackageSections retrieves an individual user resource
func (uc PackageControllerImpl) GetPackageSections(c *gin.Context) {
	packageName := c.Params.ByName("packagename")
	var resultPackage Package
	_ = uc.session.DB("gunde5dkcom").C("packages").Find(bson.M{"pname": packageName}).One(&resultPackage)

	var sectionResults []bson.M
	for _, section := range resultPackage.Sections {
		sectionObjectID := section["ObjectID"].(string)
		oid := bson.ObjectIdHex(sectionObjectID)

		var resultSection bson.M
		_ = uc.session.DB("gunde5dkcom").C("sections").Find(bson.M{"_id": oid}).One(&resultSection)

		sectionResults = append(sectionResults, resultSection)
	}
	c.JSON(http.StatusOK, sectionResults)
}
