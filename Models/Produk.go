package Models

type Produk struct {
	ID               int             `gorm:"column:id_produk;primaryKey;autoIncrement" json:"id_produk"`
	NamaProduk       string          `gorm:"column:nama_produk;varchar(255)"           json:"nama_produk"`
	HargaProduk      float64         `gorm:"column:harga_produk"                       json:"harga_produk"`
	StokProduk		 int 	         `gorm:"column:stok_produk"                        json:"stok_produk"`
	IDKategoriProduk int             `gorm:"column:id_kategori_produk"                 json:"id_kategori_produk"`
	KategoriProduk   KategoriProduk  `gorm:"foreignKey:IDKategoriProduk"               json:"kategori_produk"`
}

type ProdukInput struct {
	NamaProduk        string  `json:"nama_produk"`
	HargaProduk       float64 `json:"harga_produk"`
	StokProduk		  int 	  `json:"stok_produk"`
	IDKategoriProduk  int     `json:"id_kategori_produk"`
}

type ProdukResponse struct {
	ID                int     `json:"id_produk"`
	NamaProduk        string  `json:"nama_produk"`
	HargaProduk       float64 `json:"harga_produk"`
	StokProduk		  int 	  `json:"stok_produk"`
	IDKategoriProduk  int     `json:"id_kategori_produk"`
}

type ProdukListResponse struct {
	Produks []ProdukResponse `json:"produks"`
}