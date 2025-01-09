// middleware/auth.go
package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

const JWTSecretKey = "twoj-tajny-klucz-jwt"

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Pobieramy token z nagłówka
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Brak tokenu autoryzacji", http.StatusUnauthorized)
			return
		}

		// Rozdzielamy "Bearer" od tokenu
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
			http.Error(w, "Nieprawidłowy format tokenu", http.StatusUnauthorized)
			return
		}

		// Weryfikujemy token
		token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
			return []byte(JWTSecretKey), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Nieprawidłowy token", http.StatusUnauthorized)
			return
		}

		// Wyciągamy ID użytkownika z tokenu
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Nieprawidłowy format tokenu", http.StatusUnauthorized)
			return
		}

		userID := uint(claims["user_id"].(float64))

		// Dodajemy ID do kontekstu żądania
		ctx := context.WithValue(r.Context(), "userID", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
