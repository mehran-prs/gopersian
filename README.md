# GoPersian

A Go Lang package for dealing with Persian [text shaping](https://www.unicode.org/reports/tr9/#Shaping).  
You can use this package to print shaped persian/arabic texts on images and pdf files.

__Note__: Inspired by [goarabic](https://github.com/01walid/goarabic) but more simple and clear for persian/arabic
languages.

## Installation

```text
go get github.com/mehran-prs/gopersian
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/mehran-prs/gopersian"
)

func main() {
	fmt.Println("سلام")
	fmt.Println(gopersian.Shape("سلام"))
	fmt.Println(gopersian.Reorder(gopersian.Shape("سلام")))
	fmt.Println(gopersian.RTL("سلام"))
}
```

## Documentation

You can see docs on [gopkgdoc](https://godoc.org/github.com/mehran-prs/gopersian).

## Contributing

Contributions are greatly appreciated. Please fork this repository, make your changes, and open a pull request.

## License

MIT License.
