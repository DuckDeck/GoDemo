package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Goods struct {
	gorm.Model
	ShopID     string    //商家ID
	UnitID     int       //商品编码
	CatID      int       //分类ID
	UnitTitle  string    //商品标题
	SaleDesc   string    //销售信息
	SalNum     uint      //销售数量
	CollectNum uint      //收藏数
	CheckNum   uint      //点击数
	Stratus    int       //	状态
	Keyword    string    //关键字
	CreateTime time.Time // 创建时间
}

//对于商品的图文，视频 等专门用表来记录
type GoodDesc struct {
	ID          int    //ID
	UnitID      int    //	商品ID
	DescType    int    //	商品描述类型 1 文字 2 图片 3 视频
	DescContent string //描述内容
	DescName    string //	描述名称，一般指图片名或者视频 名
	Sort        int    //	排序
}

//商品分类
type GoodsCat struct {
	ID       int    //
	PID      int    //
	Order    int    //
	Name     string //
	FullPath string //
	Level    int    //
}

type Sku struct {
	ID     int
	UnitID int
	Price  float64
	Size   string
	Color  string
	Look   string
}

type UnitSize struct {
	gorm.Model
	cat  GoodsCat
	XXXS float32
	XXS  float32
	XS   float32
	S    float32
	M    float32
	L    float32
	XL   float32
	XXL  float32
	XXXL float32
}
