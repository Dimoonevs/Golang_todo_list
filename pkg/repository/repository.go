package repository

import (
	"github.com/jmoiron/sqlx"
	"todo-list-project"
)

type Authorization interface {
	CreateUser(user todo_list_project.User) (int, error)
	GetUser(username, password string) (todo_list_project.User, error)
}

type TodoList interface {
	Create(userId int, list todo_list_project.TodoList) (int, error)
	GetAll(userId int) ([]todo_list_project.TodoList, error)
	GetById(userId int, listId int) (todo_list_project.TodoList, error)
	Delete(userId int, listId int) error
	Update(userId, listId int, input todo_list_project.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, list todo_list_project.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todo_list_project.TodoItem, error)
	GetById(userId, itemId int) (todo_list_project.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input todo_list_project.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
