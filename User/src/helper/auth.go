package helper

import (
	"enhanced-notes/src/domain"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)
type Auth struct {
	Secret string
}

func SetupAuth(Secret string) Auth {
	return Auth{
		Secret: Secret,
	}
}


func (auth Auth) CreateHashedPassword(password string) (string, error) {

	if len(password) < 6 {
		return "", errors.New("password length should be at least 6 characters long")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		// log actual error and report to logging tool
		return "", errors.New("password hash failed")
	}

	return string(hashPassword), nil
}

func (auth Auth) GenerateToken(id uint64, email string) (string, error) {

	if id == 0 || email == "" {
		return "", errors.New("required inputs are missing to generate token")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"email":   email,
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(auth.Secret))

	if err != nil {
		return "", errors.New("unable to signed the token")
	}

	return tokenStr, nil
}


func (auth Auth) VerifyPassword(plainPassword string, hashedPassword string) error {

	if len(plainPassword) < 6 {
		return errors.New("password length should be at least 6 characters long")
	}

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))

	if err != nil {
		return errors.New("password is incorrect")
	}

	return nil
}


func (auth Auth) VerifyToken(jwtToken string) (domain.User, error) {
	tokenArr := strings.Split(jwtToken, " ")
	if len(tokenArr) != 2 {
		return domain.User{}, nil
	}
	tokenStr := tokenArr[1]
	if tokenArr[0] != "Bearer" {
		return domain.User{}, errors.New("invalid token")
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unknown signing method %v", token.Header)
		}
		return []byte(auth.Secret), nil
	})

	if err != nil {
		return domain.User{}, errors.New("invalid signing method")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return domain.User{}, errors.New("token is expired")
		}

		user := domain.User{}
		user.ID = uint64(claims["user_id"].(float64))
		user.Email = claims["email"].(string)
		return user, nil
	}

	return domain.User{}, errors.New("token verification failed")
}
func (auth Auth) Authorize(ctx fiber.Ctx) error {
	authHeader := ctx.GetReqHeaders()["Authorization"]

	user, err := auth.VerifyToken(authHeader[0])

	if err == nil && user.ID > 0 {
		ctx.Locals("user", user)
		return ctx.Next()
	} else {
		return ctx.Status(401).JSON(&fiber.Map{
			"message": "authorization failed",
			"reason":  err,
		})
	}

}

func (auth Auth) GetCurrentUser(ctx fiber.Ctx) domain.User {
	user := ctx.Locals("user")
	return user.(domain.User)

}