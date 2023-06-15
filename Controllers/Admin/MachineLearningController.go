package Admin

import (
	"fmt"
	"math"
	"sort"
	"net/http"
	"SmartStockPrediction/Models"
	"SmartStockPrediction/Database"
	"SmartStockPrediction/Utils"
)

type RecommendedProduct struct {
	Name       string  `json:"nama_produk"`
	Presentase float64 `json:"presentase"`
	Stok       int     `json:"stok_produk"`
}

// ------------------- Pendekatan Berdasarkan Jumlah Stok Ter 'SEDIKIT' -------------------
func GetRecommendation(w http.ResponseWriter, r *http.Request) {
	recommendations := GetProductRecommendations()

	// Mengurutkan rekomendasi berdasarkan presentase secara descending
	recommendedProducts := make([]RecommendedProduct, 0)
	for namaProduk, presentase := range recommendations {
		var produk Models.Produk
		if err := Database.DB.Where("nama_produk = ?", namaProduk).First(&produk).Error; err != nil {
			Utils.Logger(2, fmt.Sprintf("Admin/MachineLearning.go -> GetRecommendation() - %s", err.Error()))
			continue
		}

		recommendedProducts = append(recommendedProducts, RecommendedProduct{
			Name:       namaProduk,
			Presentase: presentase,
			Stok:       produk.StokProduk,
		})
	}
	sort.Slice(recommendedProducts, func(i, j int) bool {
		if recommendedProducts[i].Presentase == recommendedProducts[j].Presentase {
			return recommendedProducts[i].Stok < recommendedProducts[j].Stok
		}
		return recommendedProducts[i].Presentase > recommendedProducts[j].Presentase
	})

	response := make(map[string]interface{})
	if len(recommendedProducts) > 0 {
		highestRecommended := recommendedProducts[0]
		response["Rekomendasi Utama"] = map[string]float64{
			highestRecommended.Name: roundFloat(highestRecommended.Presentase, 3),
		}

		otherRecommendations := make(map[string]float64)
		for _, product := range recommendedProducts[1:] {
			otherRecommendations[product.Name] = roundFloat(product.Presentase, 3)
		}
		response["Rekomendasi Lain"] = otherRecommendations
	}

	Utils.ResponseJSON(w, http.StatusOK, response)
	Utils.Logger(3, "Admin/MachineLearning.go -> GetRecommendation()")
	return
}

func GetProductRecommendations() map[string]float64 {
	var produkList []Models.Produk
	if err := Database.DB.Find(&produkList).Error; err != nil {
		Utils.Logger(2, fmt.Sprintf("Admin/MachineLearning.go -> GetProductRecommendations() - %s", err.Error()))
		return nil
	}

	var detailTransaksiList []Models.DetailTransaksi
	if err := Database.DB.Find(&detailTransaksiList).Error; err != nil {
		Utils.Logger(2, fmt.Sprintf("Admin/MachineLearning.go -> GetProductRecommendations() - %s", err.Error()))
		return nil
	}

	// Menghitung presentase penjualan berdasarkan jumlah produk terjual dan stok barang
	recommendations := make(map[string]float64)
	for _, dt := range detailTransaksiList {
		var produk Models.Produk
		if err := Database.DB.Where("id_produk = ?", dt.IDProduk).First(&produk).Error; err != nil {
			// fmt.Println("Gagal mengambil data produk dengan ID:", dt.IDProduk)
			Utils.Logger(2, fmt.Sprintf("Admin/MachineLearning.go -> GetProductRecommendations() - %s", err.Error()))
			continue
		}

		jumlahTerjual := float64(dt.JumlahProduk)
		stokBarang := float64(produk.StokProduk)

		presentase := (jumlahTerjual / stokBarang) * 100
		recommendations[produk.NamaProduk] = presentase
	}

	return recommendations
}

func roundFloat(value float64, precision int) float64 {
	rounding := math.Pow(10, float64(precision))
	return math.Round(value*rounding) / rounding
}

// ------------------- Pendekatan Berdasarkan Jumlah Terjual Ter 'BANYAK' -------------------
func GetRecommendationByTransaction(w http.ResponseWriter, r *http.Request) {
	recommendations := GetProductRecommendationsByTransaction()

	// Mengurutkan rekomendasi berdasarkan presentase penjualan secara descending
	recommendedProducts := make([]RecommendedProduct, 0)
	for namaProduk, presentase := range recommendations {
		var produk Models.Produk
		if err := Database.DB.Where("nama_produk = ?", namaProduk).First(&produk).Error; err != nil {
			// fmt.Println("Gagal mengambil data produk:", err)
			Utils.Logger(2, fmt.Sprintf("Admin/MachineLearning.go -> GetRecommendationByTransaction() - %s", err.Error()))
			continue
		}

		recommendedProducts = append(recommendedProducts, RecommendedProduct{
			Name:       namaProduk,
			Presentase: presentase,
			Stok:       produk.StokProduk,
		})
	}
	sort.Slice(recommendedProducts, func(i, j int) bool {
		if recommendedProducts[i].Presentase == recommendedProducts[j].Presentase {
			// Jika presentase sama, urutkan berdasarkan jumlah penjualan (terbalik)
			return recommendedProducts[i].Stok > recommendedProducts[j].Stok
		}
		return recommendedProducts[i].Presentase > recommendedProducts[j].Presentase
	})

	response := make(map[string]interface{})
	if len(recommendedProducts) > 0 {
		highestRecommended := recommendedProducts[0]
		response["Rekomendasi Utama"] = map[string]float64{
			highestRecommended.Name: roundFloat(highestRecommended.Presentase, 3),
		}

		otherRecommendations := make(map[string]float64)
		for _, product := range recommendedProducts[1:] {
			otherRecommendations[product.Name] = roundFloat(product.Presentase, 3)
		}
		response["Rekomendasi Lain"] = otherRecommendations
	}

	Utils.ResponseJSON(w, http.StatusOK, response)
	Utils.Logger(3, "Admin/MachineLearning.go -> GetRecommendationByTransaction()")
	return
}

func GetProductRecommendationsByTransaction() map[string]float64 {
	var produkList []Models.Produk
	if err := Database.DB.Find(&produkList).Error; err != nil {
		fmt.Println("Gagal mengambil data produk:", err)
		Utils.Logger(2, fmt.Sprintf("Admin/MachineLearning.go -> GetProductRecommendationsByTransaction() - %s", err.Error()))
		return nil
	}

	var detailTransaksiList []Models.DetailTransaksi
	if err := Database.DB.Find(&detailTransaksiList).Error; err != nil {
		fmt.Println("Gagal mengambil data detail transaksi:", err)
		Utils.Logger(2, fmt.Sprintf("Admin/MachineLearning.go -> GetProductRecommendationsByTransaction() - %s", err.Error()))
		return nil
	}

	// Menghitung presentase penjualan berdasarkan jumlah produk terjual dan stok barang
	recommendations := make(map[string]float64)
	for _, dt := range detailTransaksiList {
		var produk Models.Produk
		if err := Database.DB.Where("id_produk = ?", dt.IDProduk).First(&produk).Error; err != nil {
			fmt.Println("Gagal mengambil data produk dengan ID:", dt.IDProduk)
			continue
		}

		jumlahTerjual := float64(dt.JumlahProduk)
		stokBarang := float64(produk.StokProduk)

		presentase := (jumlahTerjual / stokBarang) * 100
		recommendations[produk.NamaProduk] = presentase
	}

	return recommendations
}