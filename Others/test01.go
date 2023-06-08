// package Controllers 

// import (
// 	"strconv"
// 	"net/http"
// 	"encoding/json"
// 	"github.com/gorilla/mux"
// 	"SmartStockPrediction/Models"
// )

// func GetRecommendation(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	idProduk, err := strconv.Atoi(params["id"])

// 	if err != nil {
// 		http.Error(w, "request tidak diterima", http.StatusBadRequest)
// 		return
// 	}

// 	var detailTransaksis Models.DetailTransaksi

// 	db.Where("id_produk = ?", idProduk).Find(&detailTransaksis)

// 	var xData [][]float64
// 	var yData []float64
	
// 	for _, detailTransaksi := range detailTransaksis {
// 		xData = append(xData, []float64{float64(detailTransaksi.JumlahProduk)})
// 		yData = append(yData, detailTransaksi.HargaProduk)
// 	}

// 	xMat := mat.NewDense(len(xData), len(xData[0]), flattenMatrix(xData))
// 	yVec := mat.NewVecDense(len(yData), yData)

// 	// Train linear regression model
// 	model := regression.Linear{}
// 	model.Fit(xMat, yVec)

// 	stokProduk := model.Predict([]float64{float64(10)})

// 	response := struct {
// 		Rekomendasi int `json:"rekomendasi"`
// 	}{int(stokProduk)}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)
// }
