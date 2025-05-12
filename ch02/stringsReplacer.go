package main

import (
	"fmt"
	"strings"
)

func main() {
	badString := "GZ rZcks!"
	replacer := strings.NewReplacer("Z", "o")
	goodString := replacer.Replace(badString)
	fmt.Println(goodString)
}
