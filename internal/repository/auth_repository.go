package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/rizalarfiyan/be-tilik-jalan/constant"
	"github.com/rizalarfiyan/be-tilik-jalan/internal/model"
	"github.com/rizalarfiyan/be-tilik-jalan/logger"
	"github.com/rs/zerolog"
)

type AuthRepository interface {
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	UpdateLastLoggedIn(ctx context.Context, id int) error
	Insert(ctx context.Context, user model.NewUser) (*uuid.UUID, error)
}

type authRepository struct {
	db  *sql.DB
	log *zerolog.Logger
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{
		db:  db,
		log: logger.Get("auth_repository"),
	}
}

func (r *authRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	var res model.User
	query := `SELECT id, email, name, role, is_active, last_logged_in, created_at, updated_at FROM users WHERE email = $1`
	err := r.db.QueryRowContext(ctx, query, email).Scan(&res.Id, &res.Email, &res.Name, &res.Role, &res.IsActive, &res.LastLoggedIn, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *authRepository) UpdateLastLoggedIn(ctx context.Context, id int) error {
	query := `UPDATE users SET last_logged_in = NOW() WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *authRepository) Insert(ctx context.Context, user model.NewUser) (*uuid.UUID, error) {
	var id uuid.UUID
	query := `INSERT INTO users (email, name, role, is_active, last_logged_in, created_at, updated_at) VALUES ($1, $2, $3, true, NOW(), NOW(), NOW()) RETURNING id`
	err := r.db.QueryRowContext(ctx, query, user.Email, user.Name, constant.AuthGuest).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &id, nil
}
