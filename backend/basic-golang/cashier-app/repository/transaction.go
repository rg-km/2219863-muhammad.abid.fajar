package repository

type TransactionRepository struct {
	cartItemRepository CartItemRepository
}

func NewTransactionRepository(cartItemRepository CartItemRepository) TransactionRepository {
	return TransactionRepository{cartItemRepository}
}

func (u *TransactionRepository) Pay(amount int) (int, error) {
<<<<<<< HEAD
	t, err := u.cartItemRepository.TotalPrice()
	if err != nil {
		return 0, err
	}
	amount -= t
	return amount, nil
=======
	return 0, nil // TODO: replace this
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
}
