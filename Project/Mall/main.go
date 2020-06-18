package main

import (
	"GoDemo/Project/Mall/core"
)

func main() {

	//initialize.Mysql()
	//defer global.G_DB.Close()
	//initialize.DBTables()

	// tool.GetSubID(602)
	// tool.GetSubID(606)
	// tool.GetSubID(603)
	// tool.GetSubID(605)

	core.RunServer()
}
