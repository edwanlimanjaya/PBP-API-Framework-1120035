package model

type User struct {
	Id_user   int    `json : id_user`
	Name_user string `json : name_user`
	Email     string `json : email`
	Password  string `json : password`
}
