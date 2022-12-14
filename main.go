package main

import (
	"fmt"
	"golang_basic/Go_Todo_App/app/models"
)

func main() {
	/*
		fmt.Println(config.Config.Port)
		fmt.Println(config.Config.SQLDriver)
		fmt.Println(config.Config.DbName)
		fmt.Println(config.Config.LogFile)

		log.Println("test")
	*/
	fmt.Println(models.Db)

	/*
		u := &models.User{}
		u.Name = "Test"
		u.Email = "test@exampl.com"
		u.PassWord = "testtest"
		fmt.Println(u)

		u.CreateUser()
	*/
	u, _ := models.GetUser(1)
	fmt.Println(u)

	u.Name = "Test2"
	u.Email = "test2@example.com"
	u.UpdateUser()
	u, _ = models.GetUser(1)
	fmt.Println(u)

	u.DeleteUser()
	u, _ = models.GetUser(1)
	fmt.Println(u)
}
