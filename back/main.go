package main

import (
	"github.com/akiradomi/workspace/go-practice/back/config"
	"github.com/akiradomi/workspace/go-practice/back/controller"
	"github.com/akiradomi/workspace/go-practice/back/db"
	"github.com/akiradomi/workspace/go-practice/back/repository"
	"github.com/akiradomi/workspace/go-practice/back/router"
	"github.com/akiradomi/workspace/go-practice/back/usecase"
	"github.com/akiradomi/workspace/go-practice/back/utils"
	"github.com/akiradomi/workspace/go-practice/back/validator"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//ログ設定
	utils.LoggingSettings(config.Config.Logging)
	//DBインスタンス
	db := db.NewDB()
	taskValidator := validator.NewTaskValidator()
	userValidator := validator.NewUserValidator()
	//user_repositoryのインスタンス化
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)
	//user_usecaseのインスタンス化
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	tasskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	//user_controllerのインスタンス化
	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(tasskUsecase)
	//routerのインスタンス化
	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))
}
