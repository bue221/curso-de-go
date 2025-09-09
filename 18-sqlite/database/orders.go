package database

import "database/sql"

type OrderRepository struct {
	Db *sql.DB
}

type Order struct {	
	Id int
	Product string
	Amount float64
}

func (r *OrderRepository) CreateTable() error {
	_, err := r.Db.Exec(`CREATE TABLE IF NOT EXISTS orders (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	product TEXT,
	amount REAL
	);
	`)

	return err
}

func (r *OrderRepository) CreateOrder(order Order) error {
	_, err := r.Db.Exec("INSERT INTO orders (product, amount) VALUES (?, ?)", order.Product, order.Amount)
	return err
}

// GetAllOrders obtiene todas las órdenes
func (r *OrderRepository) GetAllOrders() ([]Order, error) {
	rows, err := r.Db.Query("SELECT id, product, amount FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []Order
	for rows.Next() {
		var order Order
		err := rows.Scan(&order.Id, &order.Product, &order.Amount)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

// GetOrderByID obtiene una orden por su ID
func (r *OrderRepository) GetOrderByID(id int) (Order, error) {
	var order Order
	err := r.Db.QueryRow("SELECT id, product, amount FROM orders WHERE id = ?", id).
		Scan(&order.Id, &order.Product, &order.Amount)
	
	if err != nil {
		return order, err
	}
	
	return order, nil
}

// UpdateOrder actualiza una orden existente
func (r *OrderRepository) UpdateOrder(order Order) error {
	result, err := r.Db.Exec("UPDATE orders SET product = ?, amount = ? WHERE id = ?", 
		order.Product, order.Amount, order.Id)
	
	if err != nil {
		return err
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	
	return nil
}

// DeleteOrder elimina una orden por ID
func (r *OrderRepository) DeleteOrder(id int) error {
	result, err := r.Db.Exec("DELETE FROM orders WHERE id = ?", id)
	
	if err != nil {
		return err
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	
	return nil
}

// GetOrdersByProduct obtiene órdenes filtradas por producto
func (r *OrderRepository) GetOrdersByProduct(product string) ([]Order, error) {
	rows, err := r.Db.Query("SELECT id, product, amount FROM orders WHERE product LIKE ?", "%"+product+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []Order
	for rows.Next() {
		var order Order
		err := rows.Scan(&order.Id, &order.Product, &order.Amount)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

// GetOrdersByAmountRange obtiene órdenes dentro de un rango de precios
func (r *OrderRepository) GetOrdersByAmountRange(minAmount, maxAmount float64) ([]Order, error) {
	rows, err := r.Db.Query("SELECT id, product, amount FROM orders WHERE amount BETWEEN ? AND ?", 
		minAmount, maxAmount)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []Order
	for rows.Next() {
		var order Order
		err := rows.Scan(&order.Id, &order.Product, &order.Amount)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

// CountOrders cuenta el total de órdenes
func (r *OrderRepository) CountOrders() (int, error) {
	var count int
	err := r.Db.QueryRow("SELECT COUNT(*) FROM orders").Scan(&count)
	return count, err
}

