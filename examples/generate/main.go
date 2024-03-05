package main

import (
	"fmt"

	"github.com/moroz/uuidv7-go"
)

func main() {
	uuid := uuidv7.Generate()
	fmt.Println(uuid)
}
