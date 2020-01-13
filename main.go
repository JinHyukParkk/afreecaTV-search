package main

import (
	"github.com/JinHyukParkk/randomProject/route"
)

func main() {
	router := route.Init()
	router.Logger.Debug(router.Start(":8080"))
}

