package api

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No Authorization header"})
			return
		}

		// Убираем "Bearer "
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		tokenString = strings.TrimSpace(tokenString) // Убираем лишние пробелы

		// ПЕЧАТАЕМ ТОКЕН В КОНСОЛЬ (для проверки)
		fmt.Println("Token received:", tokenString)
		fmt.Println("Secret used:", os.Getenv("JWT_SECRET"))
		// ПЕЧАТАЕМ ТОКЕН В КОНСОЛЬ (для отладки)
		fmt.Println("--- DEBUG TOKEN ---")
		fmt.Printf("Received token: [%s]\n", tokenString)
		fmt.Printf("Token Length: %d\n", len(tokenString))
		fmt.Println("-------------------")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Проверяем метод подписи
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			// ПЕЧАТАЕМ ОШИБКУ В КОНСОЛЬ
			fmt.Println("JWT Parse Error:", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: " + err.Error()})
			return
		}
		
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if val, ok := claims["user_id"].(float64); ok {
        		userID := uint(val)
        		c.Set("user_id", userID) // Теперь в контексте лежит uint
    		}
			c.Next()
		} else {
			fmt.Println("Token claims invalid")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		}
	}
}