package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecret"

// GenerateToken creates a new JWT token with user details.
func GenerateToken(email, role, userId string) (string, error) {
	// Define token claims
	claims := jwt.MapClaims{
		"Email": email,
		"ID":    userId,
		"Role":  role,
		"exp":   time.Now().Add(time.Hour * 2).Unix(), // Token expiration time
	}

	// Create a new token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and return the token
	return token.SignedString([]byte(secretKey))
}

// VerifyToken parses and validates the JWT token.
func VerifyToken(tokenString string) (string, string, string, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate token signing method
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

	// Extract claims from the token
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", tokenString, errors.New("invalid token claims")
	}

	// Extract user details from claims
	email, emailOk := claims["Email"].(string)
	userId, userIdOk := claims["ID"].(string)
	role, roleOk := claims["Role"].(string)

	if !emailOk || !userIdOk || !roleOk {
		return "", "", tokenString, errors.New("missing claims")
	}

	return email, role, userId, nil
}

// ПРИКЛАД Як витягати інфу з токену:
// func SomeProtectedHandler(w http.ResponseWriter, r *http.Request) {
//     // Отримання інформації з контексту
//     userId, ok := r.Context().Value("userId").(string)
//     if !ok {
//         http.Error(w, "User ID not found in context", http.StatusInternalServerError)
//         return
//     }

//     // Виконання логіки з отриманим userId
//     w.Write([]byte("User ID: " + userId))
// }