package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecret"

func GenerateToken(email, role, userId string) (string, error) {
	claims := jwt.MapClaims{
		"Email": email,
		"ID":    userId,
		"Role":  role,
		"exp":   time.Now().Add(time.Hour * 2).Unix(), // Token expiration time
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(tokenString string) (string, string, string, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return "", "", tokenString, errors.New("could not parse token")
	}

	if !parsedToken.Valid {
		return "", "", tokenString, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", tokenString, errors.New("invalid token claims")
	}

	exp, ok := claims["exp"].(float64)
	if !ok || time.Unix(int64(exp), 0).Before(time.Now()) {
		return "", "", "", errors.New("token expired")
	}

	email, emailOk := claims["Email"].(string)
	userId, userIdOk := claims["ID"].(string)
	role, roleOk := claims["Role"].(string)

	if !emailOk || !userIdOk || !roleOk {
		return "", "", tokenString, errors.New("missing claims")
	}

	return email, role, userId, nil
}
