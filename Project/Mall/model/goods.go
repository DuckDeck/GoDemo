package model

import (
	"GoDemo/Project/Mall/global"
	"time"

	"github.com/jinzhu/gorm"
)

type Goods struct {
	gorm.Model

	UnitID int //商品编码 生成方法  分类加标题hash

	UnitTitle  string    //商品标题
	SaleDesc   string    //销售信息
	SalNum     int       //销售数量
	CollectNum int       //收藏数
	CheckNum   int       //点击数
	Status     int       //	状态   1上架 2 下架
	Keyword    string    //关键字
	CreateTime time.Time // 创建时间
	Price      float64
	OrgPrice   float64
	Link       string //链接到原先的地方
	MainImage  string
	//one to one
	GoodDesc GoodDesc
	//one to Many
	SkuS []Sku

	GoodsCat   GoodsCat
	GoodsCatID int
	Merchant   Merchant
	MerchantID int
}

//对于商品的图文，视频 等专门用表来记录
type GoodDesc struct {
	ID          uint   //ID
	GoodsID     int    `gorm:"index"` //	商品ID
	DescType    int    //	商品描述类型 1 文字 2 图片 3 视频
	DescContent string `gorm:"type:text(5000)"` //描述内容
	DescName    string //	描述名称，一般指图片名或者视频 名
	Sort        int    //	排序
}

//商品分类
type GoodsCat struct {
	ID       int //ID
	SubCatID int
	PID      int    //上级分类ID
	Sort     int    // 排序
	Name     string // 名称
	FullPath string // 全路径
	Level    int    // 分类等级
	Logo     string // 图标
	Link     string
}

type Sku struct {
	ID      int     //ID
	GoodsID int     `gorm:"index"` //商品ID
	Price   float64 //价格
	Size    string  //尺 只对于特定商品
	Color   string  //色 只对于特定商品
	Look    string  //样式
}

//对于衣服的size
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

func GetCats(cat int) (goodsCats []GoodsCat, err error) {
	global.G_DB.Where("pid = ?", cat).Find(&goodsCats)
	return
}
