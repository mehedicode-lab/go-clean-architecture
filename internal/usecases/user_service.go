package usecases

import (
	"errors"

	"github.com/mehedicode-lab/go-clean-architecture/internal/domain"
	"github.com/mehedicode-lab/go-clean-architecture/pkg/security"
)

type UserService struct {
	Repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

// Register new user
func (s *UserService) Register(fullName, email, password string) error {
	hashedPassword, err := security.HashPassword(password)
	if err != nil {
		return err
	}

	user := &domain.User{
		FullName: fullName,
		Email:    email,
		Password: hashedPassword,
	}

	return s.Repo.CreateUser(user)
}

// Login User and return both Access Token and Refresh Token
func (s *UserService) Login(email, password string) (string, string, *domain.User, error) {
	// Fetch user from the repository by email
	user, err := s.Repo.GetUserByEmail(email)
	if err != nil {
		return "", "", nil, errors.New("invalid credentials")
	}

	// Check if the provided password matches the stored hash
	if !security.CheckPasswordHash(password, user.Password) {
		return "", "", nil, errors.New("invalid credentials")
	}

	// Generate both access and refresh tokens
	accessToken, refreshToken, err := security.GenerateTokens(user.Email)
	if err != nil {
		return "", "", nil, err
	}

	// Return both tokens along with user info
	return accessToken, refreshToken, user, nil
}

// Get user by email
func (s *UserService) GetUserByEmail(email string) (*domain.User, error) {
	return s.Repo.GetUserByEmail(email)
}
