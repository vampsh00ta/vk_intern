package service

import "context"

func (s service) Login(ctx context.Context, username string) (string, error) {
	ctxTx := s.repo.Begin(ctx)
	customer, err := s.repo.GetCustomerByUsername(ctxTx, username)
	if err != nil {
		return "", nil
	}
	jwtToken, err := createToken(customer)
	if err != nil {
		return "", nil
	}
	return jwtToken, nil
}
func (s service) Verify(ctx context.Context, token string) error {

	if err := verifyToken(token); err != nil {
		return err
	}
	return nil
}
func (s service) IsAdmin(ctx context.Context, token string) error {

	if err := verifyToken(token); err != nil {
		return err
	}
	return nil
}
