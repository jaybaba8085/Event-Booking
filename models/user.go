package models

import (
	"errors"
	"fmt"
	"log"

	"example.com/rest-api/db"
	"example.com/rest-api/utils"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

type User struct {
	ID       int    `binding:"-"`
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `
        INSERT INTO users(email, password) 
        VALUES(?, ?)
    `

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return fmt.Errorf("error executing SQL statement: %v", err)
	}

	UserId, err := result.LastInsertId()

	u.ID = int(UserId)
	return err
}

func (u *User) ValidateCredential() error {

	query := "SELECT  id, email, password FROM  users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievehHashPassword string
	err := row.Scan(&u.ID, &u.Email, &retrievehHashPassword)

	logger.Info("Checking user credentials")

	if err != nil {
		return errors.New("invalid credentials. Please try again")
	}

	log.Println(retrievehHashPassword)

	passwordIValid := utils.CheckPasswordHash(u.Password, retrievehHashPassword)

	if !passwordIValid {
		return errors.New("invalid credentials. Password does not match")
	}

	return nil
}

func GetAllUsers() ([]User, error) {
	var users []User

	query := "SELECT id, email, password FROM users"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Email, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
