package initialize

import (
	"GoDemo/Project/Mall/global"
	"GoDemo/Project/Mall/model"
)

func DBTables() {
	db := global.G_DB
	db.SingularTable(true)
	db.AutoMigrate(model.Banner{})
	db.AutoMigrate(model.Recommend{})
	global.G_LOG.Debug("register table success")
}
