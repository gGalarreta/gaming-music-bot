package main

import (
	"github.com/subosito/gotenv"
	"bot/app/routes"
)

func init() {
	gotenv.Load()
}

func main()  {
	routes.HandleRequests()
}
