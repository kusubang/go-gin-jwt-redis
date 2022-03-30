package main

import (
	"github.com/kusubang/auth/internal/routes"
)

func main() {

	r := routes.SetupRouter()

	r.Run()

}
