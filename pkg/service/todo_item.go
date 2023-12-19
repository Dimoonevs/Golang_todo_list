package service

import (
	todo_list_project "todo-list-project"
	"todo-list-project/pkg/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (i *TodoItemService) Create(userId, listId int, list todo_list_project.TodoItem) (int, error) {
	_, err := i.listRepo.GetById(userId, listId)
	if err != nil {
		return 0, err
	}
	return i.repo.Create(listId, list)
}

func (i *TodoItemService) GetAll(userId, listId int) ([]todo_list_project.TodoItem, error) {
	return i.repo.GetAll(userId, listId)

}
func (i *TodoItemService) GetById(userId, itemId int) (todo_list_project.TodoItem, error) {
	return i.repo.GetById(userId, itemId)
}

func (i *TodoItemService) Delete(userId, itemId int) error {
	return i.repo.Delete(userId, itemId)
}
func (i *TodoItemService) Update(userId, itemId int, input todo_list_project.UpdateItemInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return i.repo.Update(userId, itemId, input)

}
