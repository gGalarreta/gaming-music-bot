package main

import (
	"github.com/subosito/gotenv"
	"fmt"
	"bot/app/routes"
)

func init() {
	gotenv.Load()
}

func main()  {
	fmt.Println("Hello")
	routes.AppSetup()
}