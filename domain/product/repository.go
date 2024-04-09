package product

import (
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

var (
	ErrProductNotFound     = errors.New("the product was not found")
	ErrProductAlreadyExist = errors.New("the product already exists")
)

type ProductRepository interface {
	GetAll() ([]Product, error)
	GetByID(uuid.UUID) (Product, error)
	Add(Product) error
	Update(Product) error
	Delete(uuid.UUID) error
}

type MySQLProductRepository struct {
	db *sql.DB
}

func NewMySQLProductRepository(db *sql.DB) *MySQLProductRepository {
	return &MySQLProductRepository{db: db}
}

func (r *MySQLProductRepository) GetAll() ([]Product, error) {
	rows, err := r.db.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.item.ID, &product.item.Name, &product.price)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *MySQLProductRepository) GetByID(id uuid.UUID) (Product, error) {
	var product Product
	err := r.db.QueryRow("SELECT id, name, price FROM products WHERE id = ?", id.String()).Scan(&product.item.ID, &product.item.Name, &product.price)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Product{}, ErrProductNotFound
		}
		return Product{}, err
	}
	return product, nil
}

func (r *MySQLProductRepository) Add(product Product) error {
	_, err := r.db.Exec("INSERT INTO products (id, name, price) VALUES (?, ?, ?)", product.item.ID.String(), product.item.Name, product.price)
	if err != nil {
		if isDuplicateKeyError(err) {
			return ErrProductAlreadyExist
		}
		return err
	}
	return nil
}

func (r *MySQLProductRepository) Update(product Product) error {
	result, err := r.db.Exec("UPDATE products SET name = ?, price = ? WHERE id = ?", product.item.Name, product.price, product.item.ID.String())
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrProductNotFound
	}
	return nil
}

func (r *MySQLProductRepository) Delete(id uuid.UUID) error {
	result, err := r.db.Exec("DELETE FROM products WHERE id = ?", id.String())
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrProductNotFound
	}
	return nil
}

func isDuplicateKeyError(err error) bool {
	mysqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		return false
	}
	return mysqlErr.Number == 1062
}
