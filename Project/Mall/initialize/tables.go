package initialize

import (
	"GoDemo/Project/Mall/global"
	"GoDemo/Project/Mall/model"
)

func DBTables() {
	db := global.G_DB
	db.SingularTable(true)
	if !db.HasTable(&model.Banner{}) { //检查存在不存在，其实这个不用判断也行，不会重复建表的
		db.AutoMigrate(model.Banner{})
	}

	db.AutoMigrate(model.Recommend{})
	db.AutoMigrate(model.Goods{})
	db.AutoMigrate(model.GoodsCat{})
	db.AutoMigrate(model.GoodDesc{})
	db.AutoMigrate(model.Sku{})
	db.AutoMigrate(model.Merchant{})
	global.G_LOG.Debug("register table success")
}

func addGoodsCat() {

}
