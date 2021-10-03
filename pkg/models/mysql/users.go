package mysql

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/koad7/snippet-web/pkg/models" // New import
	"golang.org/x/crypto/bcrypt"              // New import
)

type UserModel struct {
	DB *sql.DB
}

// We'll use the Insert method to add a new record to the users table.
func (m *UserModel) Insert(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	stmt := `INSERT INTO users (name, email, hashed_password, created) VALUES(?, ?, ?, UTC_TIMESTAMP())`
	// Use the Exec() method to insert the user details and hashed password
	// into the users table.
	_, err = m.DB.Exec(stmt, name, email, string(hashedPassword))
	if err != nil {
		// If this returns an error, we use the errors.As() function to check
		// whether the error has the type *mysql.MySQLError. If it does, the
		// error will be assigned to the mySQLError variable. We can then check
		// whether or not the error relates to our users_uc_email key by
		// checking the contents of the message string. If it does, we return
		// an ErrDuplicateEmail error.
		var mySQLError *mysql.MySQLError
		if errors.As(err, &mySQLError) {
			if mySQLError.Number == 1062 && strings.Contains(mySQLError.Message, "users_uc_email") {
				return models.ErrDuplicateEmail
			}
		}
		return err
	}

	return nil
}

// We'll use the Authenticate method to verify whether a user exists with
// the provided email address and password. This will return the relevant
// user ID if they do.
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// We'll use the Get method to fetch details for a specific user based
// on their user ID.
func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}
