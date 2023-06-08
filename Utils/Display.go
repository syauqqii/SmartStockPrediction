package Utils

import "fmt"

func Display() {
	fmt.Printf("\n   > Endpoint (Admin) :")
	fmt.Printf("\n     - /users      -> GET         |  - /pelanggan      -> GET       |  - /kategori-produk      -> GET")
	fmt.Printf("\n     - /users      -> POST        |  - /pelanggan      -> POST      |  - /kategori-produk      -> POST")
	fmt.Printf("\n     - /users/{id} -> GET         |  - /pelanggan/{id} -> GET       |  - /kategori-produk/{id} -> GET")
	fmt.Printf("\n     - /users/{id} -> PUT         |  - /pelanggan/{id} -> PUT       |  - /kategori-produk/{id} -> PUT")
	fmt.Printf("\n     - /users/{id} -> DELETE      |  - /pelanggan/{id} -> DELETE    |  - /kategori-produk/{id} -> DELETE")

	fmt.Printf("\n\n     - /produk      -> GET        |  - /keranjang      -> GET       |")
	fmt.Printf("\n     - /produk      -> POST       |  - /keranjang      -> POST      |  - /transaksi      -> GET")
	fmt.Printf("\n     - /produk/{id} -> GET        |  - /keranjang/{id} -> GET       |  - /transaksi      -> POST")
	fmt.Printf("\n     - /produk/{id} -> PUT        |  - /keranjang/{id} -> PUT       |  - /transaksi/{id} -> GET")
	fmt.Printf("\n     - /produk/{id} -> DELETE     |  - /keranjang/{id} -> DELETE    |  - /transaksi/{id} -> DELETE")

	fmt.Printf("\n\n    ==============================================================================================")

	fmt.Printf("\n\n   > Endpoint (Kasir) :")
	fmt.Printf("\n     - /pelanggan      -> GET     |")
	fmt.Printf("\n     - /pelanggan      -> POST    |")
	fmt.Printf("\n     - /pelanggan/{id} -> GET     |  - /kategori-produk      -> GET |  - /produk      -> GET")
	fmt.Printf("\n     - /pelanggan/{id} -> PUT     |  - /kategori-produk/{id} -> GET |  - /produk/{id} -> GET")

	fmt.Printf("\n\n     - /keranjang      -> GET     |                                 |  - /detail-transaksi      -> GET")
	fmt.Printf("\n     - /keranjang      -> POST    |  - /transaksi      -> GET       |  - /detail-transaksi      -> POST")
	fmt.Printf("\n     - /keranjang/{id} -> GET     |  - /transaksi      -> POST      |  - /detail-transaksi/{id} -> GET")
	fmt.Printf("\n     - /keranjang/{id} -> PUT     |  - /transaksi/{id} -> GET       |  - /detail-transaksi/{id} -> PUT")
	fmt.Printf("\n     - /keranjang/{id} -> DELETE  |  - /transaksi/{id} -> DELETE    |  - /detail-transaksi/{id} -> DELETE")
}