package main

import (
	"exploration/controllers"

	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	m := martini.Classic()

	m.Get("/users", controllers.GetAllUser)
	m.Post("/insert/user", controllers.InsertUser)
	m.Put("/update/user/:id_user", controllers.UpdateUser)
	m.Delete("/delete/user/:id_user", controllers.DeleteUser)
	m.RunOnAddr(":8000")

}
