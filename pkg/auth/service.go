package auth

import (
	"errors"
	"os"
	"sharir/pkg"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// The above code defines a Service interface with a SignUp method that takes an InUser input and
// returns a string and an error.
// @property SignUp - SignUp is a method signature of an interface named Service. It takes an input
// parameter of type InUser and returns two values - a string and an error. The purpose of this method
// is to handle user sign up functionality.
type Service interface {
	Login(email string, password string) (string, error)
	LoginPhoneOtp(phone string) (string, error)
	SignUp(in InUser) (string, error)
}

// The type Svc contains a pointer to a Repo.
// @property repo - `repo` is a pointer to a `Repo` struct. It is likely used to access and manipulate
// data stored in a database or other data storage system. The `Svc` struct may contain methods that
// use the `repo` pointer to perform CRUD (create, read, update, delete) operations
type Svc struct {
	repo *Repo
}

// The `SignUp` function is a method of the `Svc` struct that implements the `Service` interface. It
// takes an `InUser` input parameter and returns a string and an error. It first checks if a user with
// the same email already exists in the repository using the `ReadByEmail` method. If a user with the
// same email exists, it returns an error. Otherwise, it creates a new user in the repository using the
// `Create` method and generates a JWT token with the user's ID, email, and expiration time using the
// `jwt` package. Finally, it returns the signed token as a string. This function is likely used for
// user sign up and authentication.
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

// The `Login` function is a method of the `Svc` struct that implements the `Service` interface. It
// takes a `phone` and `password` input parameters and returns a string and an error. It retrieves a
// user from the repository using the `ReadByPhoneNumber` method, compares the hashed password with the
// input password using `bcrypt.CompareHashAndPassword`, creates a JWT token with the user's ID, email,
// and expiration time, and returns the signed token as a string. This function is likely used for user
// authentication using a phone number and password verification.
func (s *Svc) Login(phone string, password string) (string, error) {
	user, err := s.repo.ReadByPhoneNumber(phone)
	if err != nil {
		return "", err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}
	claims := jwt.MapClaims{
		"userid": user.ID,
		"email":  user.Email,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refresh, err := refreshToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return refresh, nil

}

// The `LoginPhoneOtp` function is a method of the `Svc` struct that implements the `Service`
// interface. It takes a `phone` string input parameter and returns a string and an error. It retrieves
// a user from the repository using the `ReadByPhoneNumber` method, creates a JWT token with the user's
// ID, email, and expiration time, and returns the signed token as a string. This function is likely
// used for user authentication using a phone number and OTP (one-time password) verification.
func (s *Svc) LoginPhoneOtp(phone string) (string, error) {
	user, err := s.repo.ReadByPhoneNumber(phone)
	if err != nil {
		return "", err
	}
	claims := jwt.MapClaims{
		"userid": user.ID,
		"email":  user.Email,
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
