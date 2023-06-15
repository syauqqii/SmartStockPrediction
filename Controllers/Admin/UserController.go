package Admin

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"SmartStockPrediction/Database"
	"SmartStockPrediction/Models"
	"SmartStockPrediction/Utils"
)

func ListUser(w http.ResponseWriter, r *http.Request) {
	var users []Models.User
	if err := Database.DB.Find(&users).Error; err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			response := map[string]string{"message": "user tidak ditemukan"}
			Utils.ResponseJSON(w, http.StatusNotFound, response)
			Utils.Logger(2, "Admin/UserController.go -> ListUser() - 1")
		case err != nil:
			response := map[string]string{"message": err.Error()}
			Utils.ResponseJSON(w, http.StatusInternalServerError, response)
			Utils.Logger(2, "Admin/UserController.go -> ListUser() - 2")
		}
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
	Utils.Logger(3, "Admin/UserController.go -> ListUser()")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var userInput Models.UserInput

	if err := Utils.DecodeJSONBody(w, r, &userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		Utils.Logger(2, "Admin/UserController.go -> CreateUser() - 1")
		return
	}

	if userInput.Username == "" || userInput.Password == "" || userInput.Role == "" {
		response := map[string]string{"message": "equest body tidak boleh kosong"}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		Utils.Logger(2, "Admin/UserController.go -> CreateUser() - 2")
		return
	}

	if userInput.Role != "admin" && userInput.Role != "kasir" {
		response := map[string]string{"message": "role tidak ada"}
		Utils.ResponseJSON(w, http.StatusConflict, response)
		Utils.Logger(2, "Admin/UserController.go -> CreateUser() - 3")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusInternalServerError, response)
		Utils.Logger(2, "Admin/UserController.go -> CreateUser() - 4")
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
		Utils.Logger(2, "Admin/UserController.go -> CreateUser() - 5")
		return
	}

	if err := Database.DB.Create(&user).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusInternalServerError, response)
		Utils.Logger(2, "Admin/UserController.go -> CreateUser() - 6")
		return
	}

	response := map[string]string{"message": "berhasil membuat user"}
	Utils.ResponseJSON(w, http.StatusCreated, response)
	Utils.Logger(3, "Admin/UserController.go -> CreateUser()")
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["id"])
	if err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		Utils.Logger(2, "Admin/UserController.go -> GetUserByID() - 1")
		return
	}

	var user Models.User
	if err := Database.DB.First(&user, userID).Error; err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			response := map[string]string{"message": "user tidak ditemukan"}
			Utils.ResponseJSON(w, http.StatusNotFound, response)
			Utils.Logger(2, "Admin/UserController.go -> GetUserByID() - 2")
		case err != nil:
			response := map[string]string{"message": "terjadi kesalahan saat mengambil data pengguna"}
			Utils.ResponseJSON(w, http.StatusInternalServerError, response)
			Utils.Logger(2, "Admin/UserController.go -> GetUserByID() - 3")
		}
		return
	}

	response := Models.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
	}
	Utils.ResponseJSON(w, http.StatusOK, response)
	Utils.Logger(3, "Admin/UserController.go -> GetUserByID()")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["id"])
	if err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		Utils.Logger(2, "Admin/UserController.go -> UpdateUser() - 3")
		return
	}

	var user Models.User
	if err := Database.DB.First(&user, userID).Error; err != nil {
		response := map[string]string{"message": "user tidak ditemukan"}
		Utils.ResponseJSON(w, http.StatusNotFound, response)
		Utils.Logger(2, "Admin/UserController.go -> UpdateUser() - 4")
		return
	}

	var userInput Models.UserInput
	if err := Utils.DecodeJSONBody(w, r, &userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		Utils.Logger(2, "Admin/UserController.go -> UpdateUser() - 5")
		return
	}

	if userInput.Username == "" || userInput.Password == "" || userInput.Role == "" {
		response := map[string]string{"message": "request body tidak boleh kosong"}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		Utils.Logger(2, "Admin/UserController.go -> UpdateUser() - 6")
		return
	}

	if userInput.Username != user.Username {
		var existingUser Models.User
		if err := Database.DB.Where("username = ?", userInput.Username).First(&existingUser).Error; err == nil {
			response := map[string]string{"message": "username sudah ada"}
			Utils.ResponseJSON(w, http.StatusConflict, response)
			Utils.Logger(2, "Admin/UserController.go -> UpdateUser() - 7")
			return
		}
	}

	if userInput.Role != "admin" && userInput.Role != "kasir" {
		response := map[string]string{"message": "role tidak ada"}
		Utils.ResponseJSON(w, http.StatusConflict, response)
		Utils.Logger(2, "Admin/UserController.go -> UpdateUser() - 8")
		return
	}

	user.Username = userInput.Username
	user.Role = userInput.Role

	if err := Database.DB.Save(&user).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusInternalServerError, response)
		Utils.Logger(2, "Admin/UserController.go -> UpdateUser() - 9")
		return
	}

	response := map[string]string{"message": "berhasil update user"}
	Utils.ResponseJSON(w, http.StatusOK, response)
	Utils.Logger(3, "Admin/UserController.go -> UpdateUser()")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["id"])
	if err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		Utils.Logger(2, "Admin/UserController.go -> DeleteUser() - 1")
		return
	}

	var user Models.User
	if err := Database.DB.First(&user, userID).Error; err != nil {
		response := map[string]string{"message": "user tidak ditemukan"}
		Utils.ResponseJSON(w, http.StatusNotFound, response)
		Utils.Logger(2, "Admin/UserController.go -> DeleteUser() - 2")
		return
	}

	if err := Database.DB.Delete(&user).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusInternalServerError, response)
		Utils.Logger(2, "Admin/UserController.go -> DeleteUser() - 3")
		return
	}

	response := map[string]string{"message": "berhasil hapus user"}
	Utils.ResponseJSON(w, http.StatusOK, response)
	Utils.Logger(3, "Admin/UserController.go -> DeleteUser()")
}