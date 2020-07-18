package mysql

import (
	"database/sql"
	"fmt"
	"github.com/AlexRipoll/enchante_technical_interview/internal/user"
	"github.com/AlexRipoll/enchante_technical_interview/pkg/errors"
)

const (
	queryFindById    = "SELECT id, username, email, password, role FROM users WHERE id=?;"
	querySave        = "INSERT INTO users (id, username, email, password, role, created_on) VALUES (?, ?, ?, ?, ?, ?);"
	queryDelete      = "DELETE FROM users WHERE id=?;"
	queryFindAll     = "SELECT id, username, email, password, role, created_on FROM users;"
	queryFindByEmail = "SELECT id, username, email, password FROM users WHERE email=?;"
)

type repository struct {
	connection *sql.DB
}

func Repository(connection *sql.DB) user.Repository {
	return &repository{connection}
}

func (r *repository) Find(id string) (*user.Account, *errors.Rest) {
	stmt, stmtErr := r.connection.Prepare(queryFindById)
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

func (r *repository) Save(account *user.Account) *errors.Rest {
	stmt, stmtErr := r.connection.Prepare(querySave)
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

func (r *repository) Delete(id string) *errors.Rest {
	stmt, stmtErr := r.connection.Prepare(queryDelete)
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

func (r *repository) FindAll() ([]user.Account, *errors.Rest) {
	stmt, stmtErr := r.connection.Prepare(queryFindAll)
	if stmtErr != nil {
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	rows, queryErr := stmt.Query()
	if queryErr != nil {
		return nil, errors.NewInternalServerError("database error")
	}

	accounts := make([]user.Account, 0)
	for rows.Next() {
		var a user.Account
		scanErr := rows.Scan(&a.Id, &a.Username, &a.Email, &a.Password, &a.Role, &a.CreatedOn)
		if scanErr != nil {
			return nil, errors.NewNotFoundError(fmt.Sprintf("failed to scan row"))
		}
		accounts = append(accounts, a)
	}

	return accounts, nil
}

func (r *repository) FindByEmail(email string) (*user.Account, *errors.Rest) {
	stmt, stmtErr := r.connection.Prepare(queryFindByEmail)
	if stmtErr != nil {
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	row := stmt.QueryRow(email)
	var a user.Account
	scanErr := row.Scan(&a.Id, &a.Username, &a.Email, &a.Password)
	if scanErr != nil {
		return nil, errors.NewNotFoundError("no account found for the given email")
	}
	return &a, nil
}
