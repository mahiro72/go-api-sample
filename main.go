package main

import (
	"fmt"

	"github.com/mahiro72/go-api-sample/controller/router"
)

func main() {
	fmt.Println("Server Started")
	router.Serve()
}
