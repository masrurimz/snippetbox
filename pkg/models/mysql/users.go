package mysql

import (
	"database/sql"

	"masrurimz/snippetbox/pkg/models"
)

// UserModel Define User struct to interact with DB
type UserModel struct {
	DB *sql.DB
}

// Insert add user to the db
func (m *UserModel) Insert(user *models.User) (int, error) {
	return 0, nil
}

// Authenticate user and identify user from DB
func (m *UserModel) Authenticate(user *models.User) (int, error) {
	return 0, nil
}

// Get user details from DB
func (m *UserModel) Get(user *models.User) (*models.User, error) {
	return nil, nil
}
