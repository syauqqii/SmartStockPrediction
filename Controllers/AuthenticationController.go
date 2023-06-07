package Controllers

import (
	"time"
	"net/http"
	"gorm.io/gorm"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"SmartStockPrediction/Utils"
	"SmartStockPrediction/Models"
	"github.com/golang-jwt/jwt/v4"
	"SmartStockPrediction/Database"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var userInput Models.User

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	defer r.Body.Close()

	var user Models.User

	if err := Database.DB.Where("username = ?", userInput.Username).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"message": "username atau password salah"}
			Utils.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		default:
			response := map[string]string{"message": err.Error()}
			Utils.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		response := map[string]string{"message": "username atau password salah"}
		Utils.ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}

	expTime := time.Now().Add(time.Minute * 1)

	claims := &Utils.JWTClaim{
		Username: user.Username,
		Role: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-jwt-mux",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenAlgo.SignedString(Utils.JWT_KEY)

	if err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})

	if user.Role == "admin" {
		response := map[string]string{"message": "berhasil login sebagai admin"}
		Utils.ResponseJSON(w, http.StatusOK, response)
		return
	} else if user.Role == "kasir" {
		response := map[string]string{"message": "berhasil login sebagai kasir"}
		Utils.ResponseJSON(w, http.StatusOK, response)
		return
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})

	response := map[string]string{"message": "berhasil keluar"}
	Utils.ResponseJSON(w, http.StatusOK, response)
}