package model

import (
	"GoDemo/Project/Mall/global"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

type FiveCode struct {
	gorm.Model
	Word        string
	PinYin      string
	FiveCode    string
	ImgCode     string
	ImgKeyboard string
}

func GetCodes(letters []string) (codes []FiveCode, err error) {
	var sql = fmt.Sprintf(`select id,word,pin_yin,five_code,img_code,img_keyboard from five_code where word in ('%s')`, strings.Join(letters, ","))
	global.G_DB.Raw(sql).Find(&codes)
	return
}

func SaveFive(code FiveCode) {
	global.G_DB.Create(&code)
}
