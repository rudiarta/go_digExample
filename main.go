package main

import (
	"fmt"

	"go.uber.org/dig"
)

type Repository interface {
	Get() string
}

func NewRepository() Repository {
	return &newRepository{name: "name"}
}

type newRepository struct {
	name string
}

func (t *newRepository) Get() string {
	return t.name
}

type Service interface {
	Backup() string
}

func NewService(repo Repository) Service {
	return &newService{repo: repo}
}

type newService struct {
	repo Repository
}

func (t *newService) Backup() string {
	return t.repo.Get() + " mantap"
}

func main() {
	c := dig.New()

	err := c.Provide(NewRepository)
	if err != nil {
		fmt.Errorf("error: ", err)
	}

	err = c.Provide(NewService)
	if err != nil {
		fmt.Errorf("error: ", err)
	}

	err = c.Invoke(func(s Service) {
		fmt.Print(s.Backup())
	})
}

type Mock struct{}

func NewMock() Repository {
	return &Mock{}
}

func (m *Mock) Get() string {
	return "mock"
}
