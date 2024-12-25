package dip

import "fmt"

// Abstraction: Database interface
type Database interface {
	Save(user string)
}

// Low-level module (MongoDB) implements Database interface
type MongoDB struct{}

func (m *MongoDB) Save(user string) {
	fmt.Println("Saving user to MongoDB:", user)
}

// Another low-level module (PostgreSQL) implementing Database interface
type PostgreSQL struct{}

func (p *PostgreSQL) Save(user string) {
	fmt.Println("Saving user to PostgreSQL:", user)
}

// High-level module (UserService) depends on the abstraction (Database)
type UserService struct {
	db Database
}

func (u *UserService) CreateUser(user string) {
	u.db.Save(user)
}
