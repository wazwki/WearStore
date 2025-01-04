package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wazwki/WearStore/user-service/internal/domain"
	"github.com/wazwki/WearStore/user-service/pkg/metrics"
)

type RepositoryInterface interface {
	Create(ctx context.Context, user *domain.User) (string, error)
	FindByID(ctx context.Context, id string) (*domain.User, error)
	FindByMail(ctx context.Context, email string) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) (*domain.User, error)
	Delete(ctx context.Context, id string) error
}

type Repository struct {
	DataBase *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) RepositoryInterface {
	return &Repository{DataBase: db}
}

func (repo *Repository) Create(ctx context.Context, user *domain.User) (string, error) {
	start := time.Now()
	var id string
	query := `
		INSERT INTO users_table (email, name, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	err := repo.DataBase.QueryRow(
		ctx,
		query,
		user.Email, user.Name, user.Password, user.CreatedAt, user.UpdatedAt,
	).Scan(&id)
	if err != nil {
		return "", err
	}

	metrics.RepositoryDuration.WithLabelValues("user-service.Create").Observe(time.Since(start).Seconds())
	return id, nil
}

func (repo *Repository) FindByID(ctx context.Context, id string) (*domain.User, error) {
	start := time.Now()
	var user domain.User
	query := `
		SELECT id, email, name, password, role, created_at, updated_at
		FROM users_table
		WHERE id = $1
	`
	err := repo.DataBase.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	metrics.RepositoryDuration.WithLabelValues("user-service.FindByID").Observe(time.Since(start).Seconds())
	return &user, nil
}

func (repo *Repository) FindByMail(ctx context.Context, email string) (*domain.User, error) {
	start := time.Now()
	var user domain.User
	query := `
		SELECT id, email, name, password, role, created_at, updated_at
		FROM users_table
		WHERE email = $1
	`
	err := repo.DataBase.QueryRow(
		ctx,
		query,
		email,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	metrics.RepositoryDuration.WithLabelValues("user-service.FindByMail").Observe(time.Since(start).Seconds())
	return &user, nil
}

func (repo *Repository) Update(ctx context.Context, user *domain.User) (*domain.User, error) {
	start := time.Now()
	queryUpdate := `
		UPDATE users_table
		SET email = $1, name = $2, password = $3, updated_at = $4
		WHERE id = $5
		RETURNING id, email, name, password, role, created_at, updated_at
	`
	row := repo.DataBase.QueryRow(
		ctx,
		queryUpdate,
		user.Email, user.Name, user.Password, user.UpdatedAt, user.ID,
	)

	updatedUser := &domain.User{}
	err := row.Scan(
		&updatedUser.ID,
		&updatedUser.Email,
		&updatedUser.Name,
		&updatedUser.Password,
		&updatedUser.Role,
		&updatedUser.CreatedAt,
		&updatedUser.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	metrics.RepositoryDuration.WithLabelValues("user-service.Update").Observe(time.Since(start).Seconds())
	return updatedUser, nil
}

func (repo *Repository) Delete(ctx context.Context, id string) error {
	start := time.Now()
	query := `DELETE FROM users_table WHERE id = $1`
	commandTag, err := repo.DataBase.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return err
	}

	metrics.RepositoryDuration.WithLabelValues("user-service.Delete").Observe(time.Since(start).Seconds())
	return nil
}
