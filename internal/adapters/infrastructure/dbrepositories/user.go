package dbrepositories

import (
	"database/sql"

	"github.com/jmoiron/sqlx"

	"github.com/ianyong/todo-backend/internal/core/domainmodels"
	"github.com/ianyong/todo-backend/internal/errors/externalerrors"
)

type UserDatabaseRepository struct {
	db *sqlx.DB
}

func NewUserDatabaseRepository(db *sqlx.DB) *UserDatabaseRepository {
	return &UserDatabaseRepository{
		db: db,
	}
}

func (r *UserDatabaseRepository) GetByEmail(email string) (*domainmodels.User, error) {
	var user domainmodels.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE email = $1", email)
	if err == sql.ErrNoRows {
		return nil, &externalerrors.LoginError{}
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserDatabaseRepository) Add(user *domainmodels.User) (*domainmodels.User, error) {
	_, err := r.db.Exec(
		"INSERT INTO users (email, hashed_password) VALUES ($1, $2)",
		user.Email,
		user.HashedPassword,
	)
	if err != nil {
		return nil, err
	}

	return r.GetByEmail(user.Email)
}
