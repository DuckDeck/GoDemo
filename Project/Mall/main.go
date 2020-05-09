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

	tool.GetSubID(602)
	tool.GetSubID(606)
	tool.GetSubID(603)
	tool.GetSubID(605)
	return
	core.RunServer()
}
