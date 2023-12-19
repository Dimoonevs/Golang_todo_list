package service

import (
	todo_list_project "todo-list-project"
	"todo-list-project/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (l *TodoListService) Create(userId int, list todo_list_project.TodoList) (int, error) {
	return l.repo.Create(userId, list)
}
func (l *TodoListService) GetAll(userId int) ([]todo_list_project.TodoList, error) {
	return l.repo.GetAll(userId)
}
func (l *TodoListService) GetById(userId, listId int) (todo_list_project.TodoList, error) {
	return l.repo.GetById(userId, listId)
}
func (l *TodoListService) Delete(userId, listId int) error {
	return l.repo.Delete(userId, listId)
}
func (l *TodoListService) Update(userId, listId int, input todo_list_project.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return l.repo.Update(userId, listId, input)
}
