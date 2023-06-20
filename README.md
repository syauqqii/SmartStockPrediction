## Smart Stock Prediction
Smart Stock Prediction adalah sebuah sistem yang dikembangkan menggunakan bahasa pemrograman Golang untuk memprediksi pergerakan stok barang berdasarkan data stok saat ini dan jumlah produk yang terjual. Sistem ini dirancang untuk digunakan oleh dua jenis pengguna, yaitu admin dan kasir.

Admin memiliki akses penuh terhadap sistem dan memiliki fitur CRUD (Create, Read, Update, Delete) untuk barang dan kategori barang. Dengan fitur ini, admin dapat menambahkan, menghapus, memperbarui, dan mengelola informasi mengenai barang yang tersedia di toko. Admin juga dapat membuat kategori barang untuk membantu dalam mengorganisasi dan mengelompokkan barang-barang tersebut.

Kasir, sebagai pengguna lain dalam sistem, memiliki akses terbatas. Kasir dapat melakukan operasi get untuk melihat daftar barang yang tersedia dalam stok. Dengan demikian, kasir dapat dengan mudah mengecek ketersediaan barang kepada pelanggan saat melakukan transaksi.

Selain itu, sistem ini juga memungkinkan pelanggan melakukan transaksi. Pelanggan dapat membuat keranjang belanja, menambahkan atau menghapus barang dari keranjang, serta melakukan operasi CRUD untuk keranjang mereka sendiri. Dengan demikian, pelanggan dapat membuat daftar belanjaan yang disesuaikan dengan kebutuhan mereka.

Sistem ini juga dilengkapi dengan fitur prediksi stok. Dengan memanfaatkan data stok saat ini dan informasi tentang jumlah produk yang terjual, sistem dapat menghasilkan prediksi tentang stok barang di masa depan. Prediksi ini memberikan wawasan berharga kepada admin dan kasir untuk mengambil keputusan yang tepat terkait kebutuhan stok dan pengelolaan persediaan.

Secara keseluruhan, Smart Stock Prediction memberikan solusi yang efektif dalam mengelola stok barang, memprediksi kebutuhan stok di masa depan, serta memberikan akses yang sesuai untuk admin, kasir, dan pelanggan. Dengan menggunakan Golang sebagai bahasa pemrograman utama, sistem ini memberikan performa yang baik, kehandalan, dan kemudahan dalam pengembangan dan pengelolaan.

## Struktur Folder
--cooming soon--

## Instalasi
```
git clone https://github.com/syauqqii/SmartStockPrediction
```
```
cd SmartStockPrediction
```
Start mysql untuk membuat database baru kemudian sesuaikan nama database nya di .env file
```
go run .
```
atau bisa juga
```
go run main.go
```
Downlaod postman lalu import [FILE INI](https://github.com/syauqqii/SmartStockPrediction/blob/main/Others/Smart%20Stock%20Prediction.postman_collection.json) ke postman untuk test API

## Desain Database
<img src="https://github.com/syauqqii/SmartStockPrediction/blob/main/Others/Screenshot%20(764).png">
