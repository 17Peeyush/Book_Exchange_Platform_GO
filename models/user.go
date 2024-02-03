package models

import (
	"errors"

	"book-exchange.com/rest/db"
)

type User struct {
	ID       int64
	Name     string 
	Email    string `binding: "required"`
	Password string `binding: "required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(name, email, password) VALUES (?,?,?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(u.Name,u.Email, u.Password)

	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()
	u.ID = userId
	return err
}

func (u *User) ValidateCredentials() error{
	query :="SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err !=nil{
		return errors.New("Credentials invalid")
	}

	passwordIsValid := retrievedPassword == u.Password
	//passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid{
		return errors.New("Credentials invalid")
	}

	return nil
}