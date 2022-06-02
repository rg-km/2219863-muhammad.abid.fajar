package repository

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

type CartItemRepository struct {
	db *sql.DB
}

func NewCartItemRepository(db *sql.DB) *CartItemRepository {
	return &CartItemRepository{db: db}
}

func (c *CartItemRepository) FetchCartItems() ([]CartItem, error) {
	var sqlStatement string
	var cartItems []CartItem

	//TODO: add sql statement here
	//HINT: join table cart_items and products

	sqlStatement = "SELECT c.id, c.product_id, c.quantity, p.product_name, p.price FROM cart_items c JOIN products p ON c.product_id = p.id"

	rows, err := c.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var cartItem CartItem
		err := rows.Scan(&cartItem.ID, &cartItem.ProductID, &cartItem.Quantity, &cartItem.ProductName, &cartItem.Price)
		if err != nil {
			return nil, err
		}
		cartItems = append(cartItems, cartItem)
	}

	return cartItems, nil

	// return []CartItem{}, nil // TODO: replace this
}

func (c *CartItemRepository) FetchCartByProductID(productID int64) (CartItem, error) {
	var cartItem CartItem
	var sqlStatement string
	//TODO : you must fetch the cart by product id
	//HINT : you can use the where statement
	sqlStatement = "SELECT id, quantity FROM cart_items WHERE product_id = ?"

	row := c.db.QueryRow(sqlStatement, productID)
	if err := row.Scan(&cartItem.ID, &cartItem.Quantity); err != nil {
		return cartItem, err
	}

	return cartItem, nil

	// return CartItem{}, nil // TODO: replace this
}

func (c *CartItemRepository) InsertCartItem(cartItem CartItem) error {
	// TODO: you must insert the cart item
	_, err := c.db.Exec("INSERT INTO cart_items (product_id, quantity) VALUES (?, ?)", cartItem.ProductID, cartItem.Quantity)
	if err != nil {
		return err
	}

	return nil // TODO: replace this
}

func (c *CartItemRepository) IncrementCartItemQuantity(cartItem CartItem) error {
	//TODO : you must update the quantity of the cart item
	//HINT : you can use the where statement
	_, err := c.db.Exec("UPDATE cart_items SET quantity = quantity+? WHERE id = ?", cartItem.Quantity, cartItem.ID)
	if err != nil {
		return err
	}

	return nil // TODO: replace this
}

func (c *CartItemRepository) ResetCartItems() error {
	//TODO : you must reset the cart items
	//HINT : you can use the delete statement

	var sqlStatement string
	sqlStatement = "DELETE FROM cart_items"

	_, err := c.db.Exec(sqlStatement)
	if err != nil {
		return errors.New("error resetting cart items")
	}

	return nil // TODO: replace this
}

func (c *CartItemRepository) TotalPrice() (int, error) {
	var sqlStatement string
	//TODO : you must calculate the total price of the cart items
	//HINT : you can use the sum statement
	sqlStatement = "SELECT SUM(p.price * c.quantity) FROM cart_items c LEFT JOIN products p ON c.product_id = p.id"

	row, err := c.db.Query(sqlStatement)
	if err != nil {
		return 0, err
	}

	defer row.Close()

	var totalPrice int
	for row.Next() {
		var price int
		err := row.Scan(&price)
		if err != nil {
			return 0, err
		}
		totalPrice = price
	}

	return totalPrice, nil
	// return 0, nil // TODO: replace this
}
