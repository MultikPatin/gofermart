package services

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"main/internal/adapters"
	"main/internal/dtos"
	"main/internal/interfaces"
	"time"
)

var (
	ErrLoginAlreadyExists        = errors.New("login already exists")
	ErrAuthCredentialsIsNotValid = errors.New("login or password is not valid")
)

func NewUsersService(r interfaces.UsersRepository) *UsersService {
	return &UsersService{
		repo:   r,
		logger: adapters.GetLogger(),
	}
}

type UsersService struct {
	repo   interfaces.UsersRepository
	logger *zap.SugaredLogger
}

func (s *UsersService) Login(ctx context.Context, credentials dtos.AuthCredentials) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	user, err := s.repo.GetByLogin(ctx, credentials.Login)
	if err != nil {
		return -1, err
	}

	if !isEqualPasswords(credentials.Password, user.Password) {
		return -1, ErrAuthCredentialsIsNotValid
	}

	return user.ID, nil
}

func (s *UsersService) Register(ctx context.Context, credentials dtos.AuthCredentials) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	hash, err := hashPassword(credentials.Password)
	if err != nil {
		return -1, err
	}
	credentials.Password = hash

	userID, err := s.repo.Add(ctx, credentials)
	if err != nil {
		return -1, err
	}

	return userID, nil
}

func isEqualPasswords(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func hashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}
