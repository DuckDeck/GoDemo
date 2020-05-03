package main

import (
	"GoDemo/Project/Mall/core"
	"GoDemo/Project/Mall/global"
	"GoDemo/Project/Mall/initialize"
)

func main() {
	initialize.Mysql()
	defer global.G_DB.Close()
	core.RunServer()
}
