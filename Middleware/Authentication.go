package Middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"SmartStockPrediction/Utils"
)

func JWTAdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Ambil token dari cookie
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				response := map[string]string{"message": "tidak diizinkan"}
				Utils.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			}
		}
		tokenString := c.Value

		// Validasi token dan ambil klaim
		claims := &Utils.JWTClaim{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return Utils.JWT_KEY, nil
		})

		if err != nil {
			v, _ := err.(*jwt.ValidationError)
			switch v.Errors {
			case jwt.ValidationErrorSignatureInvalid:
				// Token tidak valid
				response := map[string]string{"message": "tidak diizinkan"}
				Utils.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			case jwt.ValidationErrorExpired:
				response := map[string]string{"message": "tidak diizinkan, token kadaluwarsa"}
				Utils.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			default:
				response := map[string]string{"message": "tidak diizinkan"}
				Utils.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			}
		}

		if !token.Valid {
			response := map[string]string{"message": "tidak diizinkan"}
			Utils.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		}

		// Cek peran (role) user
		if claims.Role != "admin" {
			response := map[string]string{"message": "tidak diizinkan"}
			Utils.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		}

		// Lanjutkan ke handler berikutnya
		next.ServeHTTP(w, r)
	})
}

func JWTKasirMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Ambil token dari cookie
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				response := map[string]string{"message": "tidak diizinkan"}
				Utils.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			}
		}
		tokenString := c.Value

		// Validasi token dan ambil klaim
		claims := &Utils.JWTClaim{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return Utils.JWT_KEY, nil
		})

		if err != nil {
			v, _ := err.(*jwt.ValidationError)
			switch v.Errors {
			case jwt.ValidationErrorSignatureInvalid:
				// Token tidak valid
				response := map[string]string{"message": "tidak diizinkan"}
				Utils.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			case jwt.ValidationErrorExpired:
				response := map[string]string{"message": "tidak diizinkan, token kadaluwarsa"}
				Utils.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			default:
				response := map[string]string{"message": "tidak diizinkan"}
				Utils.ResponseJSON(w, http.StatusUnauthorized, response)
				return
				}
		}

		if !token.Valid {
			response := map[string]string{"message": "tidak diizinkan"}
			Utils.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		}

		// Cek peran (role) user
		if claims.Role != "kasir" {
			response := map[string]string{"message": "tidak diizinkan"}
			Utils.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		}

		// Lanjutkan ke handler berikutnya
		next.ServeHTTP(w, r)
	})
}