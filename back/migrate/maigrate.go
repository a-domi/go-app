package main

import (
	"fmt"

	"github.com/akiradomi/workspace/go-practice/back/db"
	"github.com/akiradomi/workspace/go-practice/back/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Scccesfully Migrate")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
