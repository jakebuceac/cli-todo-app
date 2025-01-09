package data

import (
	"context"
	"database/sql"
	"log"
	"time"
)

const dbTimeout = time.Second * 3

var db *sql.DB

func New(dbPool *sql.DB) Models {
	db = dbPool

	return Models{
		Task: Task{},
	}
}

type Models struct {
	Task Task
}

type Task struct {
	ID        int
	Name      string
	Created   string
	Completed bool
}

func (t *Task) Index() ([]*Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)

	defer cancel()

	query := `SELECT id, description, is_complete, created_at FROM tasks ORDER BY id`

	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []*Task

	for rows.Next() {
		var task Task

		err := rows.Scan(
			&task.ID,
			&task.Name,
			&task.Completed,
			&task.Created,
		)

		if err != nil {
			log.Println("Error scanning", err)

			return nil, err
		}

		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (t *Task) Store(task Task) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)

	defer cancel()

	statement := `INSERT INTO tasks (description, is_complete, created_at) VALUES (?, ?, ?)`

	result, err := db.ExecContext(ctx, statement, task.Name, task.Completed, time.Now())

	if err != nil {
		log.Println("Error inserting", err)

		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (t *Task) Show(taskId int64) (*Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)

	defer cancel()

	query := `SELECT id, description, is_complete, created_at FROM tasks WHERE id = ?`

	var task Task

	row := db.QueryRowContext(ctx, query, taskId)

	err := row.Scan(
		&task.ID,
		&task.Name,
		&task.Completed,
		&task.Created,
	)

	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (t *Task) Update() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)

	defer cancel()

	statement := `UPDATE tasks SET is_complete = ? WHERE id = ?`

	_, err := db.ExecContext(ctx, statement,
		t.Completed,
		t.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (t *Task) Delete() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)

	defer cancel()

	statement := `DELETE from tasks where id = ?`

	_, err := db.ExecContext(ctx, statement, t.ID)

	if err != nil {
		return err
	}

	return nil
}
