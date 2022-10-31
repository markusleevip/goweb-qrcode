package main

import (
	"goweb-qrcode/config"
	"goweb-qrcode/global"
	"goweb-qrcode/initialize"
)

func main() {
	initialize.Initialize()
	global.Initialize()
	config.RunServer()

}
