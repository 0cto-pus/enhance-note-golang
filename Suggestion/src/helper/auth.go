package helper

/* import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)
type Auth struct {
	Secret string
}

func SetupAuth(Secret string) Auth {
	return Auth{
		Secret: Secret,
	}
}

func (auth Auth) VerifyToken(jwtToken string) (uint64, string, error) {
	tokenArr := strings.Split(jwtToken, " ")
	if len(tokenArr) != 2 || tokenArr[0] != "Bearer" {
		return 0, "", errors.New("invalid token")
	}

	tokenStr := tokenArr[1]
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method %v", token.Header)
		}
		return []byte(auth.Secret), nil
	})

	if err != nil || !token.Valid {
		return 0, "", errors.New("token verification failed")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return 0, "", errors.New("token is expired")
		}

		userID := uint64(claims["user_id"].(float64))
		email := claims["email"].(string)
		return userID, email, nil
	}

	return 0, "", errors.New("token verification failed")
}

func (auth Auth) Authorize(ctx fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")
	userID, email, err := auth.VerifyToken(authHeader)

	if err == nil && userID > 0 {
		ctx.Locals("user_id", userID)
		ctx.Locals("email", email)
		return ctx.Next()
	} else {
		return ctx.Status(401).JSON(&fiber.Map{
			"message": "Auth Error",
			"reason":  err.Error(),
		})
	}
}

func (auth Auth) GetCurrentUserID(ctx fiber.Ctx) uint64 {
	userID, ok := ctx.Locals("user_id").(uint64)
	if !ok {
		return 0
	}
	return userID
} */