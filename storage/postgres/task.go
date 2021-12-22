package postgres

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"

	pb "github.com/FarrukhibnAkbar/ToDo/genproto"
)

type taskRepo struct {
	db *sqlx.DB
}

func NewTaskRepo(db *sqlx.DB) *taskRepo {
	return &taskRepo{db: db}
}

func (r *taskRepo) Create(task pb.Task) (pb.Task, error) {
	var id string
	err := r.db.QueryRow(`
		INSERT INTO tasks(id, assignee, title, summary, deadline, status)
		VALUES($1, $2, $3, $4, $5, $6) returning id`,
		task.Id,
		task.Assignee,
		task.Title,
		task.Summary,
		task.Deadline,
		task.Status,
	).Scan(&id)
	if err != nil {
		return pb.Task{}, err
	}

	task, err = r.Get(id)
	if err != nil {
		return pb.Task{}, err
	}

	return task, nil
}

func (r *taskRepo) Get(id string) (pb.Task, error) {
	var task pb.Task
	var UpdatedAt sql.NullTime
	err := r.db.QueryRow(`
		SELECT id, assignee, title, summary, deadline, status, created_at, updated_at FROM tasks
		WHERE id=$1 AND deleted_at IS NULL`, id).Scan(
		&task.Id,
		&task.Assignee,
		&task.Title,
		&task.Summary,
		&task.Deadline,
		&task.Status,
		&task.CreatedAt,
		&UpdatedAt)
	if err != nil {
		return pb.Task{}, err
	}
	task.UpdatedAt = UpdatedAt.Time.String()

	return task, nil
}

func (r *taskRepo) List(page, limit int64) ([]*pb.Task, int64, error) {
	offset := (page - 1) * limit
	rows, err := r.db.Queryx(`
				SELECT id, assignee, title, summary, deadline, status, created_at
				FROM tasks WHERE deleted_at IS NULL LIMIT $1 OFFSET $2`, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	if err = rows.Err(); err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var (
		tasks []*pb.Task
		count int64
	)
	for rows.Next() {
		var task pb.Task
		err = rows.Scan(
			&task.Id,
			&task.Assignee,
			&task.Title,
			&task.Summary,
			&task.Deadline,
			&task.Status,
			&task.CreatedAt)
		if err != nil {
			return nil, 0, err
		}
		tasks = append(tasks, &task)
	}

	err = r.db.QueryRow(`SELECT count(*) FROM tasks WHERE deleted_at IS NULL`).Scan(&count)
	if err != nil {
		return nil, 0, err
	}

	return tasks, count, nil
}

func (r *taskRepo) Update(task pb.Task) (pb.Task, error) {
	result, err := r.db.Exec(`UPDATE tasks SET assignee=$1, title=$2, summary=$3, deadline=$4, status=$5, updated_at=current_timestamp 
						WHERE id=$6 and deleted_at`,
		&task.Assignee,
		&task.Title,
		&task.Summary,
		&task.Deadline,
		&task.Status,
		&task.Id)
	if err != nil {
		return pb.Task{}, err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return pb.Task{}, sql.ErrNoRows
	}

	task, err = r.Get(task.Id)
	if err != nil {
		return pb.Task{}, err
	}

	return task, err
}

func (r *taskRepo) Delete(id string) error {
	result, err := r.db.Exec(`UPDATE tasks SET deleted_at = $1 WHERE id = $2`, time.Now(), id)
	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *taskRepo) ListOverdue(page, limit int64, timer time.Time) ([]*pb.Task, int64, error) {
	offset := (page - 1) * limit
	rows, err := r.db.Queryx(`
				SELECT id, assignee, title, summary, deadline, status 
				FROM tasks WHERE deadline >= $1 LIMIT $2 OFFSET $3`, timer, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	if err = rows.Err(); err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var (
		tasks []*pb.Task
		count int64
	)
	for rows.Next() {
		var task pb.Task

		err = rows.Scan(
			&task.Id,
			&task.Assignee,
			&task.Title,
			&task.Summary,
			&task.Deadline,
			&task.Status)
		if err != nil {
			return nil, 0, err
		}
		tasks = append(tasks, &task)
	}

	err = r.db.QueryRow(`SELECT count(*) FROM tasks WHERE deadline >= $1`, timer).Scan(&count)
	if err != nil {
		return nil, 0, err
	}

	return tasks, count, nil
}
