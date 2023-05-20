package auth

import (
	"errors"
	"os"
	"sharir/pkg"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// The above code defines a Service interface with a SignUp method that takes an InUser input and
// returns a string and an error.
// @property SignUp - SignUp is a method signature of an interface named Service. It takes an input
// parameter of type InUser and returns two values - a string and an error. The purpose of this method
// is to handle user sign up functionality.
type Service interface {
	SignUp(in InUser) (string, error)
}

// The type Svc contains a pointer to a Repo.
// @property repo - `repo` is a pointer to a `Repo` struct. It is likely used to access and manipulate
// data stored in a database or other data storage system. The `Svc` struct may contain methods that
// use the `repo` pointer to perform CRUD (create, read, update, delete) operations
type Svc struct {
	repo *Repo
}

// This is a method called `SignUp` that belongs to the `Svc` struct and implements the `Service`
// interface. It takes an `InUser` input parameter and returns a string and an error.
func (s *Svc) SignUp(in InUser) (string, error) {
	user, err := s.repo.ReadByEmail(in.Email)
	if !(err == pkg.ErrUserNotFound) && err != nil {
		return "", err
	}
	if user.Email == in.Email {
		return "", errors.New("user with id already exists")
	}
	create, err := s.repo.Create(in)
	if err != nil {
		return "", err
	}
	claims := jwt.MapClaims{
		"userid": create.ID,
		"email":  create.Email,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refresh, err := refreshToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return refresh, nil
}

// The function creates a new instance of a service with a given repository.
func NewAuthService(repo *Repo) Service {
	return &Svc{
		repo: repo,
	}
}
