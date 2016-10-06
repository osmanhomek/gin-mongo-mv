package main

import (
	mgo "gopkg.in/mgo.v2"

	//"github.com/fvbock/endless"
	"gin-mongo-mv/controllers"
	"gin-mongo-mv/config"
	"fmt"
	"github.com/gin-gonic/gin"
)

func getSession() *mgo.Session {

	c := config.NewCfg("config.ini")
	if err := c.Load() ; err != nil {
        panic(err)
    }

	mongodb_host, err := c.ReadString("mongodb_host", "")
    if err != nil {panic(err)}

	mongodb_port, err := c.ReadString("mongodb_port", "")
    if err != nil {panic(err)}   

	mongodb_dbname, err := c.ReadString("mongodb_dbname", "")
    if err != nil {panic(err)} 

	mongodb_username, err := c.ReadString("mongodb_username", "")
    if err != nil {panic(err)}    

	mongodb_password, err := c.ReadString("mongodb_password", "")
    if err != nil {panic(err)}     
    
    dbcnnstr := "mongodb://"+mongodb_username+":"+mongodb_password+"@"+mongodb_host+":"+mongodb_port+"/"+mongodb_dbname        
	s, err := mgo.Dial(dbcnnstr)
	if err != nil {
		fmt.Println(dbcnnstr)
		panic(err)
	}
	return s
}
func main() {
	// Creates a gin router with default middleware:	
	session := getSession()
	uc := controllers.NewSectionController(session)

	router := gin.Default()
	router.GET("/someGet", uc.GetSections)

	//endless.ListenAndServe(":4242", router)
	router.Run(":4242")
}
