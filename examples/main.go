package main

import (
	"fmt"

	"github.com/mehran-prs/gopersian"
)

func main() {
	fmt.Println("سلام مهران")
	fmt.Println(gopersian.Shape("سلام"))
	fmt.Println(gopersian.Reorder(gopersian.Shape("سلام")))
	fmt.Println(gopersian.Reorder(gopersian.Shape("سلام خوبی? 1234 چ خبر؟")))
	fmt.Println(gopersian.Bidi("سلام خوبی? 1234 drink water چ خبر؟"))
}
