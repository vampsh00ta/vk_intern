package service

import (
	"context"
	"time"
)

func (s service) Login(ctx context.Context, username string) (string, error) {
	ctxTx := s.repo.Begin(ctx)
	customer, err := s.repo.GetCustomerByUsername(ctxTx, username)
	if err != nil {
		return "", nil
	}
	jwtToken, err := s.CreateAccessToken(customer, int(time.Minute*60*24*30))
	if err != nil {
		return "", nil
	}
	return jwtToken, nil
}
func (s service) Auth(ctx context.Context, token string) (bool, error) {
	res, err := s.IsAuthorized(token)
	if err != nil {
		return false, err
	}
	return res, nil
}
func (s service) IsAdmin(ctx context.Context, token string) (bool, error) {
	customer, err := s.ExtractCustomerFromToken(token)

	if err != nil {
		return false, err
	}
	return customer.Admin, nil
}
