package main

import "fmt"

type UserRepo interface {
	Save(name string)
}

type MySQLRepo struct {
}

func (m *MySQLRepo) Save(name string) {
	fmt.Println("Saving to MySQL", name)
}

type PostgreSQLRepository struct{}

func (p *PostgreSQLRepository) Save(name string) {
	fmt.Println("Saving to PostgreSQL:", name)
}

type MySqlDatabase struct {
}

func (db *MySqlDatabase) Save(name string) {
	fmt.Println("Saving to MySQL", name)
}

type UserService struct {
	repo UserRepo
}

func NewUserService(repo UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}
func (u *UserService) Register(name string) {
	u.repo.Save(name)
}
func main() {
	mySQL := &MySQLRepo{}
	service := NewUserService(mySQL)
	service.Register("Nishak")

}
