package repository

type TransactionRepository struct {
	cartItemRepository CartItemRepository
}

func NewTransactionRepository(cartItemRepository CartItemRepository) TransactionRepository {
	return TransactionRepository{cartItemRepository}
}

func (u *TransactionRepository) Pay(amount int) (int, error) {
	t, err := u.cartItemRepository.TotalPrice()
	if err != nil {
		return 0, err
	}
	amount -= t
	return amount, nil // TODO: replace this
}
