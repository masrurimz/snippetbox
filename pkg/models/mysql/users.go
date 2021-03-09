package mysql

import (
	"database/sql"
	"strings"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"

	"masrurimz/snippetbox/pkg/models"
)

// UserModel Define User struct to interact with DB
type UserModel struct {
	DB *sql.DB
}

// Insert add user to the db
func (m *UserModel) Insert(user *models.UserValidator) (int, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return 0, err
	}

	stmt := `INSERT INTO users (name, email, hashed_password, created) 
	VALUE(?, ?, ?, UTC_TIMESTAMP())`

	_, err = m.DB.Exec(stmt, user.Name, user.Email, string(hashedPassword))
	if err != nil {
		if mysqErr, ok := err.(*mysql.MySQLError); ok {
			if mysqErr.Number == 1062 && strings.Contains(mysqErr.Message, "users_uc_email") {
				return 0, models.ErrDuplicatedEmail
			}
		}
	}

	return 0, err
}

// Authenticate user and identify user from DB
func (m *UserModel) Authenticate(user *models.UserValidator) (int, error) {
	var id int
	var hashedPassword []byte

	row := m.DB.QueryRow("SELECT id, hashed_password FROM users WHERE email=?", user.Email)

	if err := row.Scan(&id, &hashedPassword); err == sql.ErrNoRows {
		return 0, models.ErrInvalidCredential
	} else if err != nil {
		return 0, err
	}

	if err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(user.Password)); err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, models.ErrInvalidCredential
	} else if err != nil {
		return 0, err
	}

	return id, nil
}

// Get user details from DB
func (m *UserModel) Get(user *models.User) (*models.User, error) {
	return nil, nil
}
