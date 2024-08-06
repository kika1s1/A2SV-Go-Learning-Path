package usecase

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kika1s1/task_manager/internal/domain/models"
	"github.com/kika1s1/task_manager/internal/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
    UserRepo repository.UserRepository
    JWTSecret string
}

func NewUserUsecase(repo repository.UserRepository, jwtSecret string) *UserUsecase {
    return &UserUsecase{
        UserRepo: repo,
        JWTSecret: jwtSecret,
    }
}

func (uc *UserUsecase) Register(user *models.User) error {
    existingUser, err := uc.UserRepo.FindByUsername(user.Username)
    if err != nil {
        return err
    }
    if existingUser != nil {
        return errors.New("user already exists")
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(hashedPassword)
    return uc.UserRepo.Create(user)
}

func (uc *UserUsecase) Login(username, password string) (string, error) {
    user, err := uc.UserRepo.GetByUsername(username)
    if err != nil {
		return "", errors.New("user not found")

    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil {
        return "", errors.New("invalid credentials")
    }

    token, err := uc.generateJWT(user)
    if err != nil {
        return "", err
    }

    return token, nil
}

func (uc *UserUsecase) generateJWT(user *models.User) (string, error) {
    claims := &jwt.StandardClaims{
        ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
        Subject:   user.ID,
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(uc.JWTSecret))
}
