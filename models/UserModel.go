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

func GetUsers() ([]Schema, error) {
	var users []Schema
	db := connection.GetDB()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user Schema
		err := rows.Scan(&user.ID, &user.Nama, &user.Umur)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
