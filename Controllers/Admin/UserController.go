package Admin

import (
	"strconv"
	"net/http"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"SmartStockPrediction/Utils"
	"SmartStockPrediction/Models"
	"SmartStockPrediction/Database"
)

func ListUser(w http.ResponseWriter, r *http.Request) {
	var users []Models.User
	if err := Database.DB.Find(&users).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	userResponses := make([]Models.UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = Models.UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Role:     user.Role,
		}
	}

	response := Models.UserListResponse{Users: userResponses}
	Utils.ResponseJSON(w, http.StatusOK, response)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var userInput Models.UserInput

	if err := Utils.DecodeJSONBody(w, r, &userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	if userInput.Role != "admin" && userInput.Role != "kasir" {
		response := map[string]string{"message": "role tidak ada"}
		Utils.ResponseJSON(w, http.StatusConflict, response)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)

	if err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusInternalServerError, response)
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
		return
	}

	if err := Database.DB.Create(&user).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "berhasil membuat user"}
	Utils.ResponseJSON(w, http.StatusCreated, response)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["id"])
	if err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var user Models.User
	if err := Database.DB.First(&user, userID).Error; err != nil {
		response := map[string]string{"message": "user tidak ditemukan"}
		Utils.ResponseJSON(w, http.StatusNotFound, response)
		return
	}

	response := Models.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
	}
	Utils.ResponseJSON(w, http.StatusOK, response)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["id"])
	if err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var user Models.User

	if err := Database.DB.First(&user, userID).Error; err != nil {
		response := map[string]string{"message": "user tidak ditemukan"}
		Utils.ResponseJSON(w, http.StatusNotFound, response)
		return
	}

	var userInput Models.UserInput
	
	if err := Utils.DecodeJSONBody(w, r, &userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	if userInput.Username != user.Username {
		var existingUser Models.User
		if err := Database.DB.Where("username = ?", userInput.Username).First(&existingUser).Error; err == nil {
			response := map[string]string{"message": "username sudah ada"}
			Utils.ResponseJSON(w, http.StatusConflict, response)
			return
		}
	}

	if userInput.Role != "admin" && userInput.Role != "kasir" {
		response := map[string]string{"message": "role tidak ada"}
		Utils.ResponseJSON(w, http.StatusConflict, response)
		return
	}

	user.Username = userInput.Username
	user.Role = userInput.Role

	if err := Database.DB.Save(&user).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "berhasil update user"}
	Utils.ResponseJSON(w, http.StatusOK, response)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["id"])
	if err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var user Models.User
	if err := Database.DB.First(&user, userID).Error; err != nil {
		response := map[string]string{"message": "user tidak ditemukan"}
		Utils.ResponseJSON(w, http.StatusNotFound, response)
		return
	}

	if err := Database.DB.Delete(&user).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "berhasil hapus user"}
	Utils.ResponseJSON(w, http.StatusOK, response)
}
