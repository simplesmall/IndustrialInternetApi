package main

import (
	"IndustrialInternetApi/config"
	"IndustrialInternetApi/routers"
)

func main() {
	routers.InitServer()
	defer config.CloseDB()
}
