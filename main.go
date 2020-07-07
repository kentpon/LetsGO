package main

import (
	"fmt"
	"github.com/kentpon/LetsGO/cfg"
	"github.com/kentpon/LetsGO/routes"
)

func main() {
	r := routes.Setup()
	fmt.Println("Service setup")
	r.Run(cfg.Addr)
}
