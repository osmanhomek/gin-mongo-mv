package main

import (
	mgo "gopkg.in/mgo.v2"

	//"github.com/fvbock/endless"
	"api.gunde5dk.com/controllers"
	"github.com/gin-gonic/gin"
)

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://usr_gunde_5dk:Rrd3b7yk28rbAwZb@ds035036.mlab.com:35036/gunde5dkcom")
	if err != nil {
		panic(err)
	}
	return s
}
func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	//uc := controllers.NewSectionController(getSession())
	uc := controllers.NewSectionController(getSession())

	router.GET("/someGet", uc.GetSections)

	//endless.ListenAndServe(":4242", router)
	router.Run(":4242")
}
