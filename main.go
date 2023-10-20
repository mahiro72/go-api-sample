package main

import (
	"fmt"

	"github.com/mahiro72/go-api-sample/web"
)

func main() {
	fmt.Println("Server Started!")
	web.Serve()
}
