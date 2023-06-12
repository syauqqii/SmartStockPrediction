package Utils

import (
        "github.com/fatih/color"
        "os"
        "time"
)

var Serv = color.New(color.FgHiCyan, color.Underline)

var white  = color.New(color.FgHiWhite)
var red    = color.New(color.FgHiRed)
var green  = color.New(color.FgHiGreen)
var yellow = color.New(color.FgHiYellow)
var cyan   = color.New(color.FgHiCyan)
var fatal  = color.New(color.FgRed)

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