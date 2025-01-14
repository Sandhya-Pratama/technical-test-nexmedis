package repository

import (
	"fmt"

	"github.com/Sandhya-Pratama/technical-test-nexmedis/db"
	"github.com/Sandhya-Pratama/technical-test-nexmedis/models"
)

type CartRepository struct {
}

func (repo *CartRepository) CreateCart(cart *models.Carts) (models.Carts, error) {
	tx, err := db.DB.Begin() // Mulai transaksi
	if err != nil {
		return models.Carts{}, err
	}

	// Kurangi stok produk
	stockQuery := `UPDATE products SET stock = stock - $1 WHERE id = $2 AND stock >= $1`
	res, err := tx.Exec(stockQuery, cart.Quantity, cart.ProductID)
	if err != nil {
		tx.Rollback() // Batalkan transaksi jika terjadi kesalahan
		return models.Carts{}, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		tx.Rollback()
		return models.Carts{}, fmt.Errorf("failed to update stock, insufficient stock or product not found")
	}

	// Tambahkan ke keranjang
	cartQuery := `
        INSERT INTO carts (user_id, product_id, quantity, created_at) 
        VALUES ($1, $2, $3, NOW()) 
        RETURNING id, user_id, product_id, quantity, created_at
    `
	row := tx.QueryRow(cartQuery, cart.UserID, cart.ProductID, cart.Quantity)

	var createdCart models.Carts
	if err := row.Scan(&createdCart.ID, &createdCart.UserID, &createdCart.ProductID, &createdCart.Quantity, &createdCart.CreatedAt); err != nil {
		tx.Rollback()
		return models.Carts{}, err
	}

	// Commit transaksi
	if err := tx.Commit(); err != nil {
		return models.Carts{}, err
	}

	return createdCart, nil
}

func (repo *CartRepository) GetCartByUserID(userID int) (*models.Carts, error) {
	var cart models.Carts
	query := `SELECT id, user_id, product_id, quantity, created_at FROM carts WHERE user_id = $1`
	err := db.DB.QueryRow(query, userID).Scan(&cart.ID, &cart.UserID, &cart.ProductID, &cart.Quantity, &cart.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("could not get cart: %v", err)
	}
	return &cart, nil
}

func (repo *CartRepository) DeleteCart(userID int) error {
	query := `DELETE FROM carts WHERE user_id = $1`
	_, err := db.DB.Exec(query, userID)
	if err != nil {
		return fmt.Errorf("could not delete cart: %v", err)
	}
	return nil
}
