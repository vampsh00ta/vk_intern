package service

import (
	"context"
	"fmt"
)

type Auth interface {
	Login(ctx context.Context, username string) (string, error)
	IsLogged(ctx context.Context, token string) (bool, error)
	IsAdmin(ctx context.Context, token string) (bool, error)
}

func (s service) Login(ctx context.Context, username string) (string, error) {
	ctxTx := s.repo.Begin(ctx)
	customer, err := s.repo.GetCustomerByUsername(ctxTx, username)
	if err != nil {
		return "", nil
	}
	jwtToken, err := s.CreateAccessToken(customer, 24*30)
	if err != nil {
		return "", nil
	}
	return jwtToken, nil
}
func (s service) IsLogged(ctx context.Context, token string) (bool, error) {
	res, err := s.ExtractCustomerFromToken(token)
	if err != nil {
		return false, err
	}
	fmt.Print(res)
	return true, nil
}
func (s service) IsAdmin(ctx context.Context, token string) (bool, error) {
	customer, err := s.ExtractCustomerFromToken(token)
	if err != nil {
		return false, err
	}
	return customer.Admin, nil
}
