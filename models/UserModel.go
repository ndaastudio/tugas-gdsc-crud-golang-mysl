package UserModel

import (
	connection "crud-golang-mysql/configs"
)

type Schema struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
	Umur int    `json:"umur"`
}

func CreateUser(user *Schema) error {
	db := connection.GetDB()

	_, err := db.Exec("INSERT INTO users (nama, umur) VALUES (?, ?)", user.Nama, user.Umur)
	if err != nil {
		return err
	}

	return nil
}
