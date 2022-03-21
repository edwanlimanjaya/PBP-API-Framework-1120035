package controllers

import (
	"exploration/model"
	"log"
	"net/http"
	"strconv"

	"github.com/go-martini/martini"
)

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	query := "select * from table_user"

	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	var user model.User
	var users []model.User

	for rows.Next() {
		if err := rows.Scan(&user.Id_user, &user.Name_user, &user.Email, &user.Password); err != nil {
			errorResponseClient(w)
			return
		} else {
			users = append(users, user)
		}
	}

	successResponseUsers(w, users)
}

func InsertUser(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	err := r.ParseForm()

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	name := r.Form.Get("name_user")
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	query := "Insert into table_user(name_user, email, password) values (?,?,?)"

	_, errQuery := db.Exec(query, name, email, password)

	if errQuery != nil {
		errorResponseServer(w)
		log.Fatal(errQuery.Error())
		return
	}

	successResponse(w)
}

func UpdateUser(w http.ResponseWriter, r *http.Request, params martini.Params) {
	db := Connect()
	defer db.Close()

	err := r.ParseForm()

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	id, _ := strconv.Atoi(params["id_user"])
	name := r.Form.Get("name_user")
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	query := "UPDATE `table_user` SET `name_user`= ?,`email`= ?,`password`= ? WHERE id_user = ?"

	results, errQuery := db.Exec(query, name, email, password, id)
	temp, _ := results.RowsAffected()

	if errQuery != nil {
		errorResponseServer(w)
		log.Fatal(errQuery.Error())
		return
	}

	if temp == 0 {
		rederactionResponse(w)
		return
	}

	successResponse(w)
}

func DeleteUser(w http.ResponseWriter, r *http.Request, params martini.Params) {
	db := Connect()
	defer db.Close()

	id, _ := strconv.Atoi(params["id_user"])

	query := "DELETE FROM `table_user` WHERE id_user = ?"

	results, err := db.Exec(query, id)
	temp, _ := results.RowsAffected()

	if err != nil {
		errorResponseServer(w)
		log.Fatal(err.Error())
		return
	}

	if temp == 0 {
		rederactionResponse(w)
		return
	}

	successResponse(w)
}
