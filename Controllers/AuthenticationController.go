package Controllers

import (
	"fmt"
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
		Utils.Logger(2, "AuthenticationController.go -> Login() - 1")
		return
	}

	defer r.Body.Close()

	var user Models.User

	if err := Database.DB.Where("username = ?", userInput.Username).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"message": "username atau password salah"}
			Utils.ResponseJSON(w, http.StatusUnauthorized, response)
			Utils.Logger(2, "AuthenticationController.go -> Login() - 2")
			return
		default:
			response := map[string]string{"message": err.Error()}
			Utils.ResponseJSON(w, http.StatusInternalServerError, response)
			Utils.Logger(2, "AuthenticationController.go -> Login() - 3")
			return
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		response := map[string]string{"message": "username atau password salah"}
		Utils.ResponseJSON(w, http.StatusUnauthorized, response)
		Utils.Logger(2, "AuthenticationController.go -> Login() - 4")
		return
	}

	// expTime := time.Now().Add(time.Minute * 1)
	expTime := time.Now().Add(time.Duration(Utils.EXP_TOKEN) * time.Minute)
	formattedExpTime := expTime.Format("2006-01-02 15:04:05")

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
		Utils.Logger(2, "AuthenticationController.go -> Login() - 5")
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})

	if user.Role == "admin" {
		response := map[string]interface{}{
			"message":      "Berhasil login sebagai admin",
			"token_expired_at":   formattedExpTime,
		}
		Utils.ResponseJSON(w, http.StatusOK, response)
		Utils.Logger(3, "AuthenticationController.go -> Login() - ADMIN")
		return
	} else if user.Role == "kasir" {
		response := map[string]interface{}{
			"message":      "Berhasil login sebagai kasir",
			"token_expired_at":   formattedExpTime,
		}
		Utils.ResponseJSON(w, http.StatusOK, response)
		Utils.Logger(3, "AuthenticationController.go -> Login() - KASIR")
		return
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	var userInput Models.UserInput

	if err := Utils.DecodeJSONBody(w, r, &userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		Utils.Logger(2, "AuthenticationController.go -> Register() - 1")
		return
	}

	if userInput.Role != "admin" && userInput.Role != "kasir" {
		response := map[string]string{"message": "role tidak ada"}
		Utils.ResponseJSON(w, http.StatusConflict, response)
		Utils.Logger(2, "AuthenticationController.go -> Register() - 2")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)

	if err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusInternalServerError, response)
		Utils.Logger(2, "AuthenticationController.go -> Register() - 3")
		return
	}

	user := Models.User{
		Username: userInput.Username,
		Password: string(hashedPassword),
		Role:     userInput.Role,
	}

	var existingUser Models.User

	if err := Database.DB.Where("username = ?", userInput.Username).First(&existingUser).Error; err == nil {
		response := map[string]string{"message": "username sudah ada"}
		Utils.ResponseJSON(w, http.StatusConflict, response)
		Utils.Logger(2, "AuthenticationController.go -> Register() - 4")
		return
	}

	if err := Database.DB.Create(&user).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusInternalServerError, response)
		Utils.Logger(2, "AuthenticationController.go -> Register() - 5")
		return
	}

	response := map[string]string{"message": "berhasil membuat user"}
	Utils.ResponseJSON(w, http.StatusCreated, response)
	Utils.Logger(3, fmt.Sprintf("AuthenticationController.go -> Register() - %s", userInput.Role))
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
	Utils.Logger(3, "AuthenticationController.go -> Logout()")
}