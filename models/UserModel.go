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

func GetUserByID(id int) (Schema, error) {
	var user Schema
	db := connection.GetDB()

	err := db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.ID, &user.Nama, &user.Umur)
	if err != nil {
		return user, err
	}

	return user, nil
}

func UpdateUser(id int, user *Schema) error {
	db := connection.GetDB()

	_, err := db.Exec("UPDATE users SET nama = ?, umur = ? WHERE id = ?", user.Nama, user.Umur, id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(id int) error {
	db := connection.GetDB()

	_, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
