package mysql

import (
	"database/sql"
	"fmt"
	"github.com/AlexRipoll/enchante_technical_interview/internal/product"
	"github.com/AlexRipoll/enchante_technical_interview/pkg/errors"
)

const (
	queryFindProductById = "SELECT id, name, price, seller_id, created_on, updated_on FROM products WHERE id=?;"
	querySaveProduct     = "INSERT INTO products (id, name, price, seller_id, created_on, updated_on) VALUES (?, ?, ?, ?, ?, ?);"
	queryDeleteProduct   = "DELETE FROM products WHERE id=?;"
	queryFindAllProducts = "SELECT id, name, price, seller_id, created_on, updated_on FROM products;"
)

type productRepository struct {
	connection *sql.DB
}

func ProductRepository(connection *sql.DB) product.Repository {
	return &productRepository{connection}
}

func (r *productRepository) Find(id string) (*product.Product, *errors.Rest) {
	stmt, stmtErr := r.connection.Prepare(queryFindProductById)
	if stmtErr != nil {
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	var p product.Product
	scanErr := row.Scan(&p.Id, &p.Name, &p.Price, &p.SellerId, &p.CreatedOn, &p.UpdatedOn )
	if scanErr != nil {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no product found with id %s", id))
	}
	return &p, nil
}

func (r *productRepository) Save(p *product.Product) *errors.Rest {
	stmt, stmtErr := r.connection.Prepare(querySaveProduct)
	if stmtErr != nil {
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result, err := stmt.Exec(p.Id, p.Name, p.Price, p.SellerId, p.CreatedOn, p.CreatedOn)
	if err != nil {
		return errors.NewInternalServerError("error when trying to insert product")
	}
	if _, err = result.LastInsertId(); err != nil {
		return errors.NewInternalServerError("something went wrong when saving product")
	}
	return nil
}

func (r *productRepository) Update(product *product.Product) *errors.Rest {
	return nil
}

func (r *productRepository) Delete(id string) *errors.Rest {
	return nil
}

func (r *productRepository) FindAll() ([]product.Product, *errors.Rest) {
	return nil, nil
}