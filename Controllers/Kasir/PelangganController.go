package Kasir

import (
	"strconv"
	"net/http"
	"github.com/gorilla/mux"
	"SmartStockPrediction/Utils"
	"SmartStockPrediction/Models"
	"SmartStockPrediction/Database"
)

func ListPelanggan(w http.ResponseWriter, r *http.Request) {
	var pelanggans []Models.Pelanggan

	if err := Database.DB.Find(&pelanggans).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	pelangganResponses := make([]Models.PelangganResponse, len(pelanggans))

	for i, pelanggan := range pelanggans {
		pelangganResponses[i] = Models.PelangganResponse{
			ID:            pelanggan.ID,
			NamaPelanggan: pelanggan.NamaPelanggan,
			NomorHP:       pelanggan.NomorHP,
		}
	}

	response := Models.PelangganListResponse{Pelanggans: pelangganResponses}
	Utils.ResponseJSON(w, http.StatusOK, response)
}

func CreatePelanggan(w http.ResponseWriter, r *http.Request) {
	var pelangganInput Models.PelangganInput

	if err := Utils.DecodeJSONBody(w, r, &pelangganInput); err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var existingPelanggan Models.Pelanggan

	if err := Database.DB.Where("nama_pelanggan = ?", pelangganInput.NamaPelanggan).First(&existingPelanggan).Error; err == nil {
		response := map[string]string{"message": "nama pelanggan sudah ada"}
		Utils.ResponseJSON(w, http.StatusConflict, response)
		return
	}

	pelanggan := Models.Pelanggan{
		NamaPelanggan: pelangganInput.NamaPelanggan,
		NomorHP:       pelangganInput.NomorHP,
	}

	if err := Database.DB.Create(&pelanggan).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "berhasil menambahkan data pelanggan"}
	Utils.ResponseJSON(w, http.StatusCreated, response)
}

func GetPelangganByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pelangganID, err := strconv.Atoi(params["id"])
	if err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var pelanggan Models.Pelanggan
	
	if err := Database.DB.First(&pelanggan, pelangganID).Error; err != nil {
		response := map[string]string{"message": "pelanggan tidak ditemukan"}
		Utils.ResponseJSON(w, http.StatusNotFound, response)
		return
	}

	response := Models.PelangganResponse{
		ID:            pelanggan.ID,
		NamaPelanggan: pelanggan.NamaPelanggan,
		NomorHP:       pelanggan.NomorHP,
	}
	Utils.ResponseJSON(w, http.StatusOK, response)
}

func UpdatePelanggan(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pelangganID, err := strconv.Atoi(params["id"])
	if err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var pelanggan Models.Pelanggan
	if err := Database.DB.First(&pelanggan, pelangganID).Error; err != nil {
		response := map[string]string{"message": "Ppelanggan tidak ditemukan"}
		Utils.ResponseJSON(w, http.StatusNotFound, response)
		return
	}

	var pelangganInput Models.PelangganInput
	if err := Utils.DecodeJSONBody(w, r, &pelangganInput); err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var existingPelanggan Models.Pelanggan

	if err := Database.DB.Where("nama_pelanggan = ? AND id_pelanggan != ?", pelangganInput.NamaPelanggan, pelangganID).First(&existingPelanggan).Error; err == nil {
		response := map[string]string{"message": "nama pelanggan sudah ada"}
		Utils.ResponseJSON(w, http.StatusConflict, response)
		return
	}

	pelanggan.NamaPelanggan = pelangganInput.NamaPelanggan
	pelanggan.NomorHP = pelangganInput.NomorHP

	if err := Database.DB.Save(&pelanggan).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "berhasil update pelanggan"}
	Utils.ResponseJSON(w, http.StatusOK, response)
}

func DeletePelanggan(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pelangganID, err := strconv.Atoi(params["id"])
	if err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var pelanggan Models.Pelanggan
	if err := Database.DB.First(&pelanggan, pelangganID).Error; err != nil {
		response := map[string]string{"message": "pelanggan tidak ditemukan"}
		Utils.ResponseJSON(w, http.StatusNotFound, response)
		return
	}

	if err := Database.DB.Delete(&pelanggan).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "berhasil hapus pelanggan"}
	Utils.ResponseJSON(w, http.StatusOK, response)
}
