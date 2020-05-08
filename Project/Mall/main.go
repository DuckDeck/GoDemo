package main

import (
	"GoDemo/Project/Mall/core"
	"GoDemo/Project/Mall/global"
	"GoDemo/Project/Mall/initialize"
	"GoDemo/Project/Mall/tool"
)

func main() {

	initialize.Mysql()
	defer global.G_DB.Close()
	initialize.DBTables()
	tool.FillCat()
	return
	core.RunServer()
}
