package user

import (
	"context"
	"database/sql"
	"github.com/alexyslozada/ecommerce/infrastructure/postgres"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/alexyslozada/ecommerce/model"
)

const table = "users"

var fields = []string{
	"id",
	"email",
	"password",
	"is_admin",
	"details",
	"created_at",
	"updated_at",
}

var (
	psqlInsert = postgres.BuildSQLInsert(table, fields)
	psqlGetAll = postgres.BuildSQLSelect(table, fields)
)

type User struct {
	db *pgxpool.Pool
}

// New returns a new User storage
func New(db *pgxpool.Pool) User {
	return User{db}
}

// Create creates a model.User
func (u User) Create(m *model.User) error {
	_, err := u.db.Exec(
		context.Background(),
		psqlInsert,
		m.ID,
		m.Email,
		m.Password,
		m.IsAdmin,
		m.Details,
		m.CreatedAt,
		postgres.Int64ToNull(m.UpdatedAt),
	)
	if err != nil {
		return err
	}

	return nil
}

func (u User) GetByEmail(email string) (model.User, error) {
	query := psqlGetAll + " WHERE email = $1"
	row := u.db.QueryRow(
		context.Background(),
		query,
		email,
	)

	return u.scanRow(row, true)
}

// GetAll gets all model.Users with Fields
func (u User) GetAll() (model.Users, error) {
	rows, err := u.db.Query(
		context.Background(),
		psqlGetAll,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := model.Users{}
	for rows.Next() {
		m, err := u.scanRow(rows, false)
		if err != nil {
			return nil, err
		}

		ms = append(ms, m)
	}

	return ms, nil
}

func (u User) scanRow(s pgx.Row, withPassword bool) (model.User, error) {
	m := model.User{}

	updatedAtNull := sql.NullInt64{}

	err := s.Scan(
		&m.ID,
		&m.Email,
		&m.Password,
		&m.IsAdmin,
		&m.Details,
		&m.CreatedAt,
		&updatedAtNull,
	)
	if err != nil {
		return m, err
	}

	m.UpdatedAt = updatedAtNull.Int64

	if !withPassword {
		m.Password = ""
	}

	return m, nil
}
