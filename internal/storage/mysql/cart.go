package mysql

import (
	"database/sql"
	"github.com/AlexRipoll/enchante_technical_interview/internal/cart"
	"github.com/AlexRipoll/enchante_technical_interview/pkg/errors"
)

const (
	querySaveOrder     = "INSERT INTO orders (id, user_id, product_id, seller_id, price, quantity) VALUES (?, ?, ?, ?, ?, ?);"
)

type cartRepository struct {
	connection *sql.DB
}

func CartRepository(connection *sql.DB) cart.Repository {
	return &cartRepository{connection}
}

func (r *cartRepository) Save(o *cart.Order) *errors.Rest {
	stmt, stmtErr := r.connection.Prepare(querySaveOrder)
	if stmtErr != nil {
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	for i, _ := range o.Items {
		result, err := stmt.Exec(o.Id, o.UserId, o.Items[i].Id, o.Items[i].SellerId, o.Items[i].Price, o.Items[i].Quantity)
		if err != nil {
			return errors.NewInternalServerError("error when trying to insert order: %s")
		}
		if _, err = result.LastInsertId(); err != nil {
			return errors.NewInternalServerError("something went wrong when saving order")
		}
	}
	return nil
}

