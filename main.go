package main

import (
	"SmartStockPrediction/Route"
	"SmartStockPrediction/Utils"
	"SmartStockPrediction/Database"
)

func main() {
	Utils.ClearScreen()
	Utils.LoadEnv()

	Database.ConnectDB()
	Route.RunRoute()
}