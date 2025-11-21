package main

import (
	"fmt"
	_ "image/png"

	"github.com/gookit/color"
	"github.com/qeesung/image2ascii/convert"
)

func main() {
	color.Warn.Prompt("⚠️  Side effects include learning, networking, and fun ⚠️  \n")

	// tips message
	// color.Info.Tips("tips style message")
	// color.Warn.Tips("tips style message")

	color.Print("*****  <bg=120,35,156>MoCo</> ")
	color.Red.Print("Code")

	color.Cyan.Print("And")
	color.Yellow.Print("Cocao")
	fmt.Println("  *****\n")

	color.Info.Tips("Location: 401 N Washington St, Rockville, MD 20850\n")
	color.Info.Tips("Date: Friday December 19th, 2025\n")
	color.Info.Tips("Time: 5:30pm - 8:30pm\n")

	// Convert and print image
	convertOptions := convert.DefaultOptions
	convertOptions.FixedWidth = 120 // Adjust width as needed
	convertOptions.FixedHeight = 30 // Adjust height as needed

	converter := convert.NewImageConverter()
	fmt.Print(converter.ImageFile2ASCIIString("code&coffee_logo.png", &convertOptions))
}
