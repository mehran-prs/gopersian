package main

import (
	"fmt"

	"github.com/mehran-prs/gopersian"
)

func main() {
	//fmt.Println("سلام")
	//fmt.Println(gopersian.Shape("سلام"))
	//fmt.Println(gopersian.Reorder(gopersian.Shape("سلام")))
	//fmt.Println(gopersian.Reorder(gopersian.Shape("سلام خوبی? 1234 چ خبر؟")))
	//fmt.Println(gopersian.RTL("سلام خوبی? 1234 drink water چ خبر؟"))
	txt := "سلام"
	fmt.Println("unchanged",[]rune(txt))
	fmt.Println("shaped",[]rune(gopersian.Shape(txt)))
	v, _ := gopersian.Reorder(txt)
	fmt.Println("reordered",[]rune(v))
	v, _ = gopersian.Reorder(gopersian.Shape(txt))
	fmt.Println("reorder_shared:",[]rune(v))
}
