package initialize

import (
	"GoDemo/Project/Mall/global"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Mysql() {
	ad := global.G_CONFIG.Mysql

	if db, err := gorm.Open("mysql", ad.Username+":"+ad.Password+"@("+ad.Path+")/"+ad.Dbname+"?"+ad.Config); err != nil {
		fmt.Println(err.Error())
		global.G_LOG.Error("DEFALUT数据库启动异常")
	} else {
		global.G_DB = db
		global.G_DB.DB().SetMaxIdleConns(ad.MaxIdleConns)
		global.G_DB.DB().SetMaxOpenConns(ad.MaxOpenConns)
		global.G_DB.LogMode(ad.LogMode)
	}

}
