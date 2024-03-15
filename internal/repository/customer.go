package repository

import (
	"context"
	"vk/internal/repository/models"
)

type Customer interface {
	GetCustomerById(ctx context.Context, id int) (models.Customer, error)
	GetCustomerByUsername(ctx context.Context, username string) (models.Customer, error)
}

func (p Pg) GetCustomerById(ctx context.Context, id int) (models.Customer, error) {
	q := `select * from customer where id = $1`
	tx := p.getTx(ctx)

	var customer models.Customer
	if err := tx.Raw(q, id).Scan(&customer).Error; err != nil {
		return models.Customer{}, err
	}
	return customer, nil
}
func (p Pg) GetCustomerByUsername(ctx context.Context, username string) (models.Customer, error) {
	q := `select * from customer where username = $1`
	tx := p.getTx(ctx)

	var customer models.Customer
	if err := tx.Raw(q, username).Scan(&customer).Error; err != nil {
		return models.Customer{}, err
	}
	return customer, nil
}
