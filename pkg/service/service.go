package service

import (
	"todo-list-project"
	"todo-list-project/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo_list_project.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo_list_project.TodoList) (int, error)
	GetAll(userId int) ([]todo_list_project.TodoList, error)
	GetById(userId int, listId int) (todo_list_project.TodoList, error)
	Delete(userId int, listId int) error
	Update(userId, listId int, input todo_list_project.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, list todo_list_project.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todo_list_project.TodoItem, error)
	GetById(userId, itemId int) (todo_list_project.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input todo_list_project.UpdateItemInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
