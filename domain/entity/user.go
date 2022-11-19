package entity

import (
	"context"
	"exercise/domain"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// UserRepository represent the user's repository contract
type UserRepository interface {
	FindByEmail(c context.Context, email string) (*User, error)
	Create(c context.Context, user *User) error
}

// UserUsecase represent the user's usecase contract
type UserUsecase interface {
	Register(ctx context.Context, user *UserRegister) (string, error)
	Login(ctx context.Context, user *UserLogin) (string, error)
}

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	NoHp      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

var signature = []byte("rahasaiabanget")

type UserRegister struct {
	Name     string `json:"name,omitempty" binding:"required"`
	Email    string `json:"email,omitempty" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
}

type UserLogin struct {
	Email    string `json:"email,omitempty" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
}

func NewUser(email, name, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		Name:     name,
		Email:    email,
		Password: string(hash),
	}, nil
}

func (u *User) GenerateJWT() (string, error) {
	claims := jwt.MapClaims{
		"user_id": u.ID,
		"exp":     time.Now().Add(1 * time.Hour).Unix(),
		"iss":     "edspert",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	stringToken, err := token.SignedString(signature)
	return stringToken, err
}

func (u *User) DecryptJWT(token string) (map[string]interface{}, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, domain.ErrUnauthorized
		}
		return signature, nil
	})

	data := make(map[string]interface{})
	if err != nil {
		return data, nil
	}

	if !parsedToken.Valid {
		return data, domain.ErrUnauthorized
	}

	return parsedToken.Claims.(jwt.MapClaims), nil
}

func (u *User) CorrectPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
