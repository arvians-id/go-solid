package main

import (
	"fmt"
	"log"
)

type User struct {
	ID   int
	Name string
}

var data = []User{
	{
		ID:   1,
		Name: "Widdy",
	},
	{
		ID:   2,
		Name: "Agung",
	},
	{
		ID:   3,
		Name: "Diyo",
	},
}

type Database interface {
	FindAll() ([]User, error)
	FindByID(id int) (User, error)
}

type PostgreSQLRepository struct {
}

func NewPostgreSQLRepository() Database {
	return &PostgreSQLRepository{}
}

func (p PostgreSQLRepository) FindAll() ([]User, error) {
	fmt.Println("Executed Query Find All PostgreSQL = `SELECT * FROM users`")
	return data, nil
}

func (p PostgreSQLRepository) FindByID(id int) (User, error) {
	fmt.Println("Executed Query Find By ID PostgreSQL = `SELECT * FROM users WHERE $1`")
	for _, val := range data {
		if val.ID == id {
			return val, nil
		}
	}

	return User{}, nil
}

type MySQLRepository struct {
}

func NewMySQLRepository() Database {
	return &MySQLRepository{}
}

func (m MySQLRepository) FindAll() ([]User, error) {
	fmt.Println("Executed Query Find All MYSQL = `SELECT * FROM users`")
	return data, nil
}

func (m MySQLRepository) FindByID(id int) (User, error) {
	fmt.Println("Executed Query Find By ID MYSQL = `SELECT * FROM users WHERE ?`")
	for _, val := range data {
		if val.ID == id {
			return val, nil
		}
	}

	return User{}, nil
}

type UserService struct {
	Database
}

func NewUserService(database Database) *UserService {
	return &UserService{
		Database: database,
	}
}

func (u *UserService) FindAll() {
	log.Println(u.Database.FindAll())
}

func (u *UserService) FindByID(id int) {
	log.Println(u.Database.FindByID(id))
}

func main() {
	// You can uncomment this DB if you want to use it and comment out another DB. Keep only one DB that can be initialized
	//db := NewMySQLRepository()

	db := NewPostgreSQLRepository()
	userService := NewUserService(db)
	userService.FindAll()
	userService.FindByID(3)
}
