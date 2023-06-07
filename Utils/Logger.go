package Utils

import (
	"github.com/fatih/color"
	"os"
	"time"
)

var Serv = color.New(color.FgHiCyan, color.Underline) // => Cyan Terang + Garis Bawah

var white  = color.New(color.FgHiWhite)  // => Putih Terang
var red    = color.New(color.FgHiRed)    // => Merah Terang
var green  = color.New(color.FgHiGreen)  // => Hijau Terang
var yellow = color.New(color.FgHiYellow) // => Kuning Terang
var cyan   = color.New(color.FgHiCyan)   // => Cyan Terang
var fatal  = color.New(color.FgRed)      // => Merah

func Logger(opt int, msg string) {
	now := time.Now()

	 white.Printf(" [%s] ", yellow.SprintFunc()(now.Format("15:04:05 2006/01/02")))

	if opt == 1 {
                cyan.Printf("INFO")
        } else if opt == 2{
                red.Printf("ERROR")
        } else if opt == 3 {
                green.Printf("SUCCESS")
        } else if opt == 4 {
                fatal.Printf("FATAL")
                white.Println(": "+msg)
                os.Exit(0)
        }

	white.Println(": "+msg)
}