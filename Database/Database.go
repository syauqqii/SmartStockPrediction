package Database

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/driver/mysql"
	"SmartStockPrediction/Utils"
	"SmartStockPrediction/Models"
)

var DB *gorm.DB

func ConnectDB() {
	config := gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}

	db, err := gorm.Open(mysql.Open(Utils.DB_CONN), &config)
	
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&Models.User{},
		&Models.Pelanggan{},
		&Models.KategoriProduk{},
		&Models.Produk{},
		&Models.Keranjang{},
		&Models.Transaksi{},
		&Models.DetailTransaksi{},
	)

	DB = db
}