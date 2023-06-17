package Models

type DetailTransaksi struct {
	ID              int     	`gorm:"column:id_detail_transaksi;primaryKey;autoIncrement" json:"id_detail_transaksi"`
	IDTransaksi     int     	`gorm:"column:id_transaksi"                                 json:"id_transaksi"`
	IDProduk        int     	`gorm:"column:id_produk"                                    json:"id_produk"`
	JumlahProduk    int     	`gorm:"column:jumlah_produk"                                json:"jumlah_produk"`
	HargaProduk     float64 	`gorm:"column:harga_produk"                                 json:"harga_produk"`
	Transaksi    	Transaksi  	`gorm:"foreignKey:IDTransaksi"                              json:"transaksi"`
	Produk       	Produk     	`gorm:"foreignKey:IDProduk"                                 json:"produk"`
}

type DetailTransaksiInput struct {
	IDTransaksi     int     `json:"id_transaksi"`
	IDProduk        int     `json:"id_produk"`
	JumlahProduk    int     `json:"jumlah_produk"`
	HargaProduk     float64 `json:"harga_produk"`
}

type DetailTransaksiResponse struct {
	ID              int     `json:"id"`
	IDTransaksi     int     `json:"id_transaksi"`
	IDProduk        int     `json:"id_produk"`
	JumlahProduk    int     `json:"jumlah_produk"`
	HargaProduk     float64 `json:"harga_produk"`
}

type DetailTransaksiListResponse struct {
	DetailTransaksis []DetailTransaksiResponse `json:"detail_transaksis"`
}