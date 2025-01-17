package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/TechBowl-japan/go-stations/model"
)

// A TODOService implements CRUD of TODO entities.
type TODOService struct {
	db *sql.DB
}

// NewTODOService returns new TODOService.
func NewTODOService(db *sql.DB) *TODOService {
	return &TODOService{
		db: db,
	}
}

// CreateTODO creates a TODO on DB.
func (s *TODOService) CreateTODO(ctx context.Context, subject, description string) (*model.TODO, error) {
	const (
		insert  = `INSERT INTO todos(subject, description) VALUES(?, ?)`
		confirm = `SELECT subject, description, created_at, updated_at FROM todos WHERE id = ?`
	)
	result, err := s.db.ExecContext(ctx, insert, subject, description)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	var created, updated time.Time
	err = s.db.QueryRowContext(ctx, confirm, id).Scan(&subject, &description, &created, &updated)
	if err != nil {
		return nil, err
	}
	return &model.TODO{
		ID:          int(id),
		Subject:     subject,
		Description: description,
		CreatedAt:   created,
		UpdatedAt:   updated,
	}, nil
}

// ReadTODO reads TODOs on DB.
func (s *TODOService) ReadTODO(ctx context.Context, prevID, size int64) ([]*model.TODO, error) {
	const (
		read       = `SELECT id, subject, description, created_at, updated_at FROM todos ORDER BY id DESC LIMIT ?`
		readWithID = `SELECT id, subject, description, created_at, updated_at FROM todos WHERE id < ? ORDER BY id DESC LIMIT ?`
	)
	return nil, nil
}

// UpdateTODO updates the TODO on DB.
func (s *TODOService) UpdateTODO(ctx context.Context, id int64, subject, description string) (*model.TODO, error) {
	const (
		update  = `UPDATE todos SET subject = ?, description = ? WHERE id = ?`
		confirm = `SELECT subject, description, created_at, updated_at FROM todos WHERE id = ?`
	)
	var confirmSubject, confirmDescription string
	var created, updated time.Time
	err := s.db.QueryRowContext(ctx, confirm, id).Scan(&confirmSubject, &confirmDescription, &created, &updated)
	if err != nil {
		return nil, err
	}

	_, err = s.db.ExecContext(ctx, update, subject, description, id)
	if err != nil {
		return nil, err
	}

	if confirmSubject == subject && confirmDescription == description {
		return nil, model.Run().Error()
	}

	err = s.db.QueryRowContext(ctx, confirm, id).Scan(&subject, &description, &created, &updated)
	if err != nil {
		return nil, err
	}
	return &model.TODO{
		ID:          int(id),
		Subject:     subject,
		Description: description,
		CreatedAt:   created,
		UpdatedAt:   updated,
	}, nil
}

// DeleteTODO deletes TODOs on DB by ids.
func (s *TODOService) DeleteTODO(ctx context.Context, ids []int64) error {
	const deleteFmt = `DELETE FROM todos WHERE id IN (?%s)`

	return nil
}
