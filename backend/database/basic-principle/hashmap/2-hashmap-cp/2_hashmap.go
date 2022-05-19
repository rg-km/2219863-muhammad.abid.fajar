package main

import "errors"

type PrimaryKey int

type InvoiceRow struct {
	ID               PrimaryKey
	SubscriptionCode string
	CustomerName     string
	Address          string
	Phone            string
}

type InvoiceDB struct {
	m map[PrimaryKey]InvoiceRow
}

func NewInvoice() *InvoiceDB {
	return &InvoiceDB{
		m: make(map[PrimaryKey]InvoiceRow),
	}
}

func (db *InvoiceDB) Insert(code string, name string, address string, phone string) {
	db.m[PrimaryKey(len(db.m)+1)] = InvoiceRow{
		ID:               PrimaryKey(len(db.m)) + 1,
		SubscriptionCode: code,
		CustomerName:     name,
		Address:          address,
		Phone:            phone,
	}
}

func (db *InvoiceDB) Where(id PrimaryKey) *InvoiceRow {
	// return InvoiceRow{} // TODO: replace this
	f := (*db).m[id]

	return &InvoiceRow{
		ID:               id,
		SubscriptionCode: f.SubscriptionCode,
		CustomerName:     f.CustomerName,
		Address:          f.Address,
		Phone:            f.Phone,
	}

}

func (db *InvoiceDB) Update(id PrimaryKey, code string, name string, address string, phone string) (*InvoiceRow, error) {
	c := db.Where(id)

	if c == nil {
		return nil, errors.New("not found")
	}

	return &InvoiceRow{
		SubscriptionCode: code,
		CustomerName:     name,
		Address:          address,
		Phone:            phone,
	}, nil

	// return nil, nil // TODO: replace this
}
