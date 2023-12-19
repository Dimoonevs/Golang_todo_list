package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
	todo_list_project "todo-list-project"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (l *TodoListPostgres) Create(userId int, list todo_list_project.TodoList) (int, error) {
	tx, err := l.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	createTodoList := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1 , $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(createTodoList, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUserList := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
	_, err = tx.Exec(createUserList, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}

func (l *TodoListPostgres) GetAll(userId int) ([]todo_list_project.TodoList, error) {
	var lists []todo_list_project.TodoList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1",
		todoListsTable, usersListsTable)
	err := l.db.Select(&lists, query, userId)
	fmt.Println(lists)

	return lists, err
}
func (l *TodoListPostgres) GetById(userId, listId int) (todo_list_project.TodoList, error) {
	var list todo_list_project.TodoList
	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description FROM %s tl
								INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2`,
		todoListsTable, usersListsTable)
	err := l.db.Get(&list, query, userId, listId)

	return list, err

}
func (l *TodoListPostgres) Delete(userId, listId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id=$1 AND ul.list_id=$2", todoListsTable, usersListsTable)
	_, err := l.db.Exec(query, userId, listId)
	return err
}

func (l *TodoListPostgres) Update(userId, listId int, input todo_list_project.UpdateListInput) error {
	setValue := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValue = append(setValue, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}
	if input.Description != nil {
		setValue = append(setValue, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	setQuery := strings.Join(setValue, ", ")
	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
		todoListsTable, setQuery, usersListsTable, argId, argId+1)
	args = append(args, listId, userId)
	logrus.Debugf("query: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := l.db.Exec(query, args...)
	return err
}
