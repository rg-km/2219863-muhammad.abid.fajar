package repository

import (
	"strconv"

	"github.com/ruang-guru/playground/backend/basic-golang/cashier-app/db"
)

type CartItemRepository struct {
	db db.DB
}

func NewCartItemRepository(db db.DB) CartItemRepository {
	return CartItemRepository{db}
}

func (u *CartItemRepository) LoadOrCreate() ([]CartItem, error) {
	records, err := u.db.Load("cart_items")
	if err != nil {
		records = [][]string{
			{"category", "product_name", "price", "quantity"},
		}
		if err := u.db.Save("cart_items", records); err != nil {
			return nil, err
		}
	}

	result := make([]CartItem, 0)
	for i := 1; i < len(records); i++ {
		price, err := strconv.Atoi(records[i][2])
		if err != nil {
			return nil, err
		}

		qty, err := strconv.Atoi(records[i][3])
		if err != nil {
			return nil, err
		}

		cartItem := CartItem{
			Category:    records[i][0],
			ProductName: records[i][1],
			Price:       price,
			Quantity:    qty,
		}
		result = append(result, cartItem)
	}

	return result, nil
}

func (u *CartItemRepository) Save(cartItems []CartItem) error {
	records := [][]string{
		{"category", "product_name", "price", "quantity"},
	}
	for i := 0; i < len(cartItems); i++ {
		records = append(records, []string{
			cartItems[i].Category,
			cartItems[i].ProductName,
			strconv.Itoa(cartItems[i].Price),
			strconv.Itoa(cartItems[i].Quantity),
		})
	}
	return u.db.Save("cart_items", records)
}

func (u *CartItemRepository) SelectAll() ([]CartItem, error) {
	cartItems, err := u.LoadOrCreate()
	if err != nil {
		return nil, err
	}
	return cartItems, nil
}

func (u *CartItemRepository) Add(product Product) error {
	cartItems, err := u.LoadOrCreate()
	if err != nil {
		return err
	}
	flag := false
	for i := 0; i < len(cartItems); i++ {
		if cartItems[i].ProductName == product.ProductName {
			flag = true
			cartItems[i].Quantity++
			return u.Save(cartItems)
		}
	}
	if flag == false {
		cartItems = append(cartItems, CartItem{
			Category:    product.Category,
			ProductName: product.ProductName,
			Price:       product.Price,
			Quantity:    1,
		})
	}
	return u.Save(cartItems) // TODO: replace this
}

func (u *CartItemRepository) ResetCartItems() error {
	cartItems, err := u.LoadOrCreate()
	if err != nil {
		return err
	}
	cartItems = nil

	return u.Save(cartItems) // TODO: replace this
}

func (u *CartItemRepository) TotalPrice() (int, error) {
	cartItems, err := u.LoadOrCreate()
	if err != nil {
		return 0, err
	}
	var dummy []int
	var total int
	for i := 0; i < len(cartItems); i++ {
		price := cartItems[i].Price
		dummy = append(dummy, int(price))
	}

	for _, v := range dummy {
		total += v
	}
	total += 2000

	return total, nil
}
