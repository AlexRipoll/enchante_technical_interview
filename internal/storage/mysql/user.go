package mysql

import (
	"database/sql"
	"fmt"
	"github.com/AlexRipoll/enchante_technical_interview/internal/user"
	"github.com/AlexRipoll/enchante_technical_interview/pkg/errors"
)

const (
	queryFindUserById    = "SELECT id, username, email, password, role FROM users WHERE id=?;"
	querySaveUser        = "INSERT INTO users (id, username, email, password, role, created_on) VALUES (?, ?, ?, ?, ?, ?);"
	queryDeleteUser      = "DELETE FROM users WHERE id=?;"
	queryFindUserByEmail = "SELECT id, username, email, password, role FROM users WHERE email=?;"
)

type userRepository struct {
	connection *sql.DB
}

func UserRepository(connection *sql.DB) user.Repository {
	return &userRepository{connection}
}

func (r *userRepository) Find(id string) (*user.Account, *errors.Rest) {
	stmt, stmtErr := r.connection.Prepare(queryFindUserById)
	if stmtErr != nil {
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	var a user.Account
	scanErr := row.Scan(&a.Id, &a.Username, &a.Email, &a.Password, &a.Role)
	if scanErr != nil {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no account found with id %s", id))
	}
	return &a, nil
}

func (r *userRepository) Save(account *user.Account) *errors.Rest {
	stmt, stmtErr := r.connection.Prepare(querySaveUser)
	if stmtErr != nil {
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result, err := stmt.Exec(account.Id, account.Username, account.Email, account.Password, account.Role, account.CreatedOn)
	if err != nil {
		return errors.NewInternalServerError("error when trying to insert account")
	}
	if _, err = result.LastInsertId(); err != nil {
		return errors.NewInternalServerError("something went wrong when saving account")
	}
	return nil
}

func (r *userRepository) Delete(id string) *errors.Rest {
	stmt, stmtErr := r.connection.Prepare(queryDeleteUser)
	if stmtErr != nil {
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, err := stmt.Exec(id)
	if err != nil {
		return errors.NewInternalServerError("error when trying to delete account")
	}

	return nil
}

func (r *userRepository) FindByEmail(email string) (*user.Account, *errors.Rest) {
	stmt, stmtErr := r.connection.Prepare(queryFindUserByEmail)
	if stmtErr != nil {
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	row := stmt.QueryRow(email)
	var a user.Account
	scanErr := row.Scan(&a.Id, &a.Username, &a.Email, &a.Password, &a.Role)
	if scanErr != nil {
		return nil, errors.NewNotFoundError("no account found for the given email")
	}
	return &a, nil
}
