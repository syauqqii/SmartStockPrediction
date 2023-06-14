package Admin

import (
	"strconv"
	"net/http"
	"gorm.io/gorm"
	"github.com/gorilla/mux"
	"SmartStockPrediction/Utils"
	"SmartStockPrediction/Models"
	"SmartStockPrediction/Database"
)

func CreateProduk(w http.ResponseWriter, r *http.Request) {
	var produkInput Models.ProdukInput

	if err := Utils.DecodeJSONBody(w, r, &produkInput); err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		Utils.Logger(2, "Admin/ProdukController.go -> CreateProduk() - 1")
		return
	}

	var existingProduk Models.Produk
	if err := Database.DB.Where("nama_produk = ?", produkInput.NamaProduk).First(&existingProduk).Error; err == nil {
		response := map[string]string{"message": "nama produk sudah ada"}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		Utils.Logger(2, "Admin/ProdukController.go -> CreateProduk() - 2")
		return
	}

	produk := Models.Produk{
		NamaProduk:        produkInput.NamaProduk,
		HargaProduk:       produkInput.HargaProduk,
		StokProduk:        produkInput.StokProduk,
		IDKategoriProduk:  produkInput.IDKategoriProduk,
	}

	if err := Database.DB.Create(&produk).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusInternalServerError, response)
		Utils.Logger(2, "Admin/ProdukController.go -> CreateProduk() - 3")
		return
	}

	response := map[string]string{"message": "berhasil menambahkan produk"}
	Utils.ResponseJSON(w, http.StatusCreated, response)
	Utils.Logger(3, "Admin/ProdukController.go -> CreateProduk()")
}

func GetProdukByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	produkID, err := strconv.Atoi(params["id"])
	if err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		Utils.Logger(2, "Admin/ProdukController.go -> GetProdukByID() - 1")
		return
	}

	var produk Models.Produk
	if err := Database.DB.Preload("KategoriProduk").First(&produk, produkID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			response := map[string]string{"message": "produk tidak ditemukan"}
			Utils.ResponseJSON(w, http.StatusNotFound, response)
			Utils.Logger(2, "Admin/ProdukController.go -> GetProdukByID() - 2")
			return
		}
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusInternalServerError, response)
		Utils.Logger(2, "Admin/ProdukController.go -> GetProdukByID() - 3")
		return
	}

	produkResponse := Models.ProdukResponse{
		ID:                produk.ID,
		NamaProduk:        produk.NamaProduk,
		HargaProduk:       produk.HargaProduk,
		StokProduk:        produk.StokProduk,
		IDKategoriProduk:  produk.IDKategoriProduk,
	}

	response := map[string]interface{}{"produk": produkResponse}
	Utils.ResponseJSON(w, http.StatusOK, response)
	Utils.Logger(3, "Admin/ProdukController.go -> GetProdukByID()")
}

func GetAllProduk(w http.ResponseWriter, r *http.Request) {
	var produks []Models.Produk
	if err := Database.DB.Preload("KategoriProduk").Find(&produks).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusInternalServerError, response)
		Utils.Logger(2, "Admin/ProdukController.go -> GetAllProduk() - 1")
		return
	}

	var produkResponses []Models.ProdukResponse
	for _, produk := range produks {
		produkResponse := Models.ProdukResponse{
			ID:                produk.ID,
			NamaProduk:        produk.NamaProduk,
			HargaProduk:       produk.HargaProduk,
			StokProduk:        produk.StokProduk,
			IDKategoriProduk:  produk.IDKategoriProduk,
		}
		produkResponses = append(produkResponses, produkResponse)
	}

	response := Models.ProdukListResponse{Produks: produkResponses}
	Utils.ResponseJSON(w, http.StatusOK, response)
	Utils.Logger(3, "Admin/ProdukController.go -> GetAllProduk()")
}

func UpdateProduk(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	produkID, err := strconv.Atoi(params["id"])
	if err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		Utils.Logger(2, "Admin/ProdukController.go -> UpdateProduk() - 1")
		return
	}

	var produk Models.Produk
	if err := Database.DB.First(&produk, produkID).Error; err != nil {
		response := map[string]string{"message": "produk tidak ditemukan"}
		Utils.ResponseJSON(w, http.StatusNotFound, response)
		Utils.Logger(2, "Admin/ProdukController.go -> UpdateProduk() - 2")
		return
	}

	var produkInput Models.ProdukInput
	if err := Utils.DecodeJSONBody(w, r, &produkInput); err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		Utils.Logger(2, "Admin/ProdukController.go -> UpdateProduk() - 3")
		return
	}

	var existingProduk Models.Produk
	if err := Database.DB.Where("id_produk != ?", produkID).Where("nama_produk = ?", produkInput.NamaProduk).First(&existingProduk).Error; err == nil {
		response := map[string]string{"message": "nama produk sudah ada"}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		Utils.Logger(2, "Admin/ProdukController.go -> UpdateProduk() - 4")
		return
	}

	produk.NamaProduk = produkInput.NamaProduk
	produk.HargaProduk = produkInput.HargaProduk
	produk.StokProduk = produkInput.StokProduk
	produk.IDKategoriProduk = produkInput.IDKategoriProduk

	if err := Database.DB.Save(&produk).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusInternalServerError, response)
		Utils.Logger(2, "Admin/ProdukController.go -> UpdateProduk() - 5")
		return
	}

	response := map[string]string{"message": "berhasil mengupdate produk"}
	Utils.ResponseJSON(w, http.StatusOK, response)
	Utils.Logger(3, "Admin/ProdukController.go -> UpdateProduk()")
}

func DeleteProduk(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	produkID, err := strconv.Atoi(params["id"])
	if err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		Utils.Logger(2, "Admin/ProdukController.go -> DeleteProduk() - 1")
		return
	}

	var produk Models.Produk
	if err := Database.DB.First(&produk, produkID).Error; err != nil {
		response := map[string]string{"message": "produk tidak ditemukan"}
		Utils.ResponseJSON(w, http.StatusNotFound, response)
		Utils.Logger(2, "Admin/ProdukController.go -> DeleteProduk() - 2")
		return
	}

	if err := Database.DB.Delete(&Models.Produk{}, produkID).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusInternalServerError, response)
		Utils.Logger(2, "Admin/ProdukController.go -> DeleteProduk() - 3")
		return
	}

	response := map[string]string{"message": "berhasil menghapus produk"}
	Utils.ResponseJSON(w, http.StatusOK, response)
	Utils.Logger(3, "Admin/ProdukController.go -> DeleteProduk()")
}
